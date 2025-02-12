package errors

import (
	"context"
	"encoding/json"
	"fmt"

	apm "github.com/jhonquirama/hexagonal-scaffolding/pkg/monitor/elastic-apm"
	"google.golang.org/grpc/codes"
)

type (
	CustomCode string

	Error struct {
		ctx           context.Context
		code          CustomCode
		message       string
		stacktrace    string
		traceID       interface{}
		transactionID interface{}
		spanID        interface{}
		httpCode      int
		grpcCode      codes.Code
		externalError error
		send          *bool
	}

	// swagger:model ErrorResponse
	ErrorResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func New(c context.Context, code CustomCode, opts ...Option) error {
	var (
		ctx       = apm.DetachedContext(c)
		skip      = 1
		customErr = getDefinition(code)
		options   = option{}
	)

	for _, funcOption := range opts {
		funcOption(&options)
	}

	if options.skip > 0 {
		skip = options.skip
	}

	funcName, line := apm.GetCallerName(skip)
	customErr.stacktrace = fmt.Sprintf("%s:%d", funcName, line)
	customErr.code = code
	customErr.ctx = ctx
	customErr.traceID = apm.TraceIDContext(ctx)
	customErr.transactionID = apm.TransactionIDContext(ctx)
	customErr.spanID = apm.SpanIDContext(ctx)

	fillWithOptions(ctx, options, &customErr)

	apm.CaptureError(ctx, &customErr)

	return &customErr
}

func fillWithOptions(ctx context.Context, options option, customErr *Error) {
	if options.message != "" {
		customErr.message = options.message
	}

	if options.httpCode > 0 {
		customErr.httpCode = options.httpCode
	}

	if options.externalError != nil {
		customErr.externalError = options.externalError

		tx := apm.TransactionFromContext(ctx)

		if tx != nil && tx.Transaction != nil && tx.TransactionData != nil {
			tx.SetLabel("error", customErr.externalError)
			tx.SetLabel("error.stack", customErr.stacktrace)
		}

		span := apm.SpanFromContext(ctx)
		if span != nil && span.Span != nil && span.SpanData != nil {
			span.SetLabel("error", customErr.externalError)
			span.SetLabel("stack", customErr.stacktrace)
		}
	}

	customErr.send = options.send
}

func (e *Error) Error() string {
	return e.message
}

func getError(err error) Error {
	switch err := err.(type) {
	case *Error:
		return *err
	default:
		return Error{
			message: err.Error(),
		}
	}
}

func Code(err error) string {
	return string(getError(err).code)
}

func Message(err error) string {
	return getError(err).message
}

func Stacktrace(err error) string {
	return getError(err).stacktrace
}

func HTTPCode(err error) int {
	return getError(err).httpCode
}

func GrpcCode(err error) codes.Code {
	return getError(err).grpcCode
}

func Context(err error) context.Context {
	return getError(err).ctx
}

func TraceID(err error) interface{} {
	return getError(err).traceID
}

func TransactionID(err error) interface{} {
	return getError(err).transactionID
}

func SpanID(err error) interface{} {
	return getError(err).spanID
}

func ExternalError(err error) error {
	return getError(err).externalError
}

func Send(err error) bool {
	parseErr := getError(err)
	if parseErr.send != nil {
		return *parseErr.send
	}
	return true
}

func UnmarshalJSON(ctx context.Context, data []byte, httpCode int) error {
	var response ErrorResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return New(ctx, UnknownError, WithError(err), WithMessage(string(data)))
	} else if response.Code == "" {
		return New(ctx, UnknownError, WithMessage(string(data)))
	}

	return New(ctx,
		CustomCode(response.Code),
		WithMessage(response.Message),
		withHTTPCode(httpCode),
	)
}
