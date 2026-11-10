package gin

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	"io"
)

const (
	processStartedBody   = "process started with body"
	processStartedParams = "process started with parameters"
	processFinished      = "process finished"
	requestBodyLabel     = "request_body"
	requestParamsLabel   = "request_params"
	responseBodyLabel    = "response_body"
)

type (
	bodyLogWriter struct {
		gin.ResponseWriter
		body *bytes.Buffer
	}
)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func MiddlewareTracking(c *gin.Context) {
	response := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}

	c.Writer = response

	if bodyBytes, err := io.ReadAll(c.Request.Body); err == nil {
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		body := map[string]interface{}{}
		var array []map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &body); err == nil && len(body) > 0 {
			if out, err := json.Marshal(body); err == nil {
				customLogger.Info(c.Request.Context(), processStartedBody, customLogger.WithObject(out))
			}
		} else if err := json.Unmarshal(bodyBytes, &array); err == nil && len(array) > 0 {
			if out, err := json.Marshal(array); err == nil {
				customLogger.Info(c.Request.Context(), processStartedBody, customLogger.WithObject(out))
			}
		}
	}

	if out, err := json.Marshal(c.Params); err == nil && len(c.Params) > 0 {
		customLogger.Info(c.Request.Context(), processStartedParams, customLogger.WithObject(out))
	}

	c.Next()

	customLogger.Info(c.Request.Context(), processFinished, customLogger.WithObject(response.body.String()))
}

func MiddlewareError(c *gin.Context) {
	c.Next()

	ctx := c.Request.Context()

	if len(c.Errors) == 0 {
		return
	}

	customLogger.Error(ctx, c.Errors.Last().Err)
	if errorFound(c.Errors.Last().Err, c) {
		return
	}

	errorFound(customError.New(ctx, customError.UnknownError), c)
}

func errorFound(err error, c *gin.Context) bool {
	if customError.Code(err) != "" {
		c.JSON(customError.HTTPCode(err), customError.ErrorResponse{
			Code:    customError.Code(err),
			Message: customError.Message(err),
		})
		return true
	}
	return false
}
