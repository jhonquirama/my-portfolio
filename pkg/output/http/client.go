package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
)

type Option func(r *option)

type option struct {
	basicAuth        *basicAuth
	headers          map[string]string
	body, response   interface{}
	maxRetries       int
	delayMillisecond time.Duration
}

type basicAuth struct {
	username string
	password string
}

func WithBody(body interface{}) Option {
	return func(r *option) {
		r.body = body
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(r *option) {
		r.headers = headers
	}
}

func WithBasicAuth(username, password string) Option {
	return func(r *option) {
		r.basicAuth = &basicAuth{
			username: username,
			password: password,
		}
	}
}

func WithResponse(response interface{}) Option {
	return func(r *option) {
		r.response = response
	}
}

func WithRetry(maxRetries int, delayMillisecond time.Duration) Option {
	return func(r *option) {
		r.maxRetries = maxRetries
		r.delayMillisecond = delayMillisecond
	}
}

type Client interface {
	Request(ctx context.Context, httpMethod, url string, opts ...Option) error
}

type client struct {
	timeout int
}

func NewClient(timeout int) Client {
	return &client{
		timeout: timeout,
	}
}

func (c *client) Request(ctx context.Context, httpMethod, url string, options ...Option) error { // nolint: gocognit
	span, ctx := apm.NewSpan(ctx, apm.Client)
	defer span.End()

	var (
		bodyResponse []byte
		bodyRequest  bytes.Buffer
		clientHTTP   *http.Client
		response     *http.Response
		request      *http.Request
		optional     option
		err          error
	)

	for _, optionItem := range options {
		if optionItem != nil {
			optionItem(&optional)
		}
	}

	if optional.body != nil {
		if err = json.NewEncoder(&bodyRequest).Encode(optional.body); err != nil {
			return customError.New(ctx, customError.UnknownError, customError.WithError(err))
		}
	}

	if request, err = http.NewRequestWithContext(ctx, httpMethod, url, &bodyRequest); err != nil {
		return customError.New(ctx, customError.UnknownError, customError.WithError(err))
	}

	request.Header.Set("Content-Type", "application/json")

	for key, value := range optional.headers {
		request.Header.Set(key, value)
	}

	if optional.basicAuth != nil {
		request.SetBasicAuth(optional.basicAuth.username, optional.basicAuth.password)
	}

	clientHTTP = apm.WrapClient(&http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	})

	if optional.maxRetries <= 0 {
		optional.maxRetries = 1 // garantiza q como minimo se haga un intento
	}

	for i := 1; i <= optional.maxRetries; i++ {
		response, err = clientHTTP.Do(request)
		if i == optional.maxRetries { // ultimo intento
			if err != nil {
				return customError.New(ctx, customError.UnknownError, customError.WithError(err))
			}
			break
		}
		if err != nil ||
			response.StatusCode < 200 ||
			(response.StatusCode >= 300 && response.StatusCode != http.StatusNotFound) {
			if response != nil && response.Body != nil {
				response.Body.Close()
			}
			time.Sleep(optional.delayMillisecond)
			continue
		}
		break
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNoContent, http.StatusCreated, http.StatusAccepted:
		return nil
	case http.StatusOK:
		if optional.response == nil {
			return nil
		}

		if err = json.NewDecoder(response.Body).Decode(&optional.response); err != nil {
			return customError.New(ctx, customError.UnknownError, customError.WithError(err))
		}
	case http.StatusNotFound:
		return customError.New(ctx, customError.UnknownError, customError.WithError(err))
	default:
		if bodyResponse, err = io.ReadAll(response.Body); err != nil {
			return customError.New(ctx, customError.UnknownError, customError.WithError(err))
		}

		customLogger.Info(ctx, string(bodyResponse))
		return customError.UnmarshalJSON(ctx, bodyResponse, response.StatusCode)
	}

	return nil
}
