package errors

import (
	"context"
	"fmt"

	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel"
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

func New(ctx context.Context, code CustomCode, opts ...Option) error {
	var (
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

	fillWithOptions(ctx, options, &customErr)

	return &customErr
}

func fillWithOptions(_ context.Context, options option, customErr *Error) {
	if options.message != "" {
		customErr.message = options.message
	}

	if options.httpCode > 0 {
		customErr.httpCode = options.httpCode
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
