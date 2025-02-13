package logger

import (
	"context"
	"encoding/json"
	"maps"
	"os"

	logrus "github.com/sirupsen/logrus"

	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
)

type Option func(l *option)

type option struct {
	Send   *bool
	Object interface{}
}

func WithObject(object interface{}) Option {
	return func(l *option) {
		l.Object = object
	}
}

func Send(send bool) Option {
	return func(l *option) {
		l.Send = &send
	}
}

var (
	environment string            // nolint: gochecknoglobals
	logger      = &logrus.Logger{ // nolint: gochecknoglobals
		Out:   os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
			},
		},
	}
)

func SetEnvironment(env string) {
	environment = env
}

func Error(ctx context.Context, err error, options ...Option) {
	var (
		optional option
	)

	for _, option := range options {
		option(&optional)
	}

	fields := getFields(optional.Object)
	maps.Copy(fields, getErrorFields(err))
	maps.Copy(fields, getContextFields(ctx))

	if externalError := customError.ExternalError(err); externalError != nil &&
		(optional.Send == nil || *optional.Send) &&
		customError.Send(err) {
	}

	logger.WithContext(ctx).WithFields(fields).Error(err)
}

func Info(ctx context.Context, message string, options ...Option) {
	var (
		optional option
	)

	for _, option := range options {
		option(&optional)
	}

	fields := getFields(optional.Object)
	fields["message"] = message

	logger.WithContext(ctx).WithFields(fields).Info(message)
}

func Warn(ctx context.Context, err error, options ...Option) {
	var (
		optional option
	)

	for _, option := range options {
		option(&optional)
	}

	fields := getFields(optional.Object)

	if externalError := customError.ExternalError(err); externalError != nil &&
		(optional.Send == nil || *optional.Send) &&
		customError.Send(err) {
	}

	logger.WithContext(ctx).WithFields(fields).Warn(err)
}

func Fatal(ctx context.Context, err error, options ...Option) {
	var (
		optional option
	)

	for _, option := range options {
		option(&optional)
	}

	fields := getFields(optional.Object)

	if externalError := customError.ExternalError(err); externalError != nil &&
		(optional.Send == nil || *optional.Send) &&
		customError.Send(err) {
	}

	logger.WithContext(ctx).WithFields(fields).Fatal(err)
}

func getErrorFields(err error) logrus.Fields {
	fields := logrus.Fields{}
	if err == nil {
		return fields
	}

	errorStack := customError.Stacktrace(err)
	if errorStack != "" {
		fields["error.stack"] = errorStack
	}

	errorTraceID := customError.TraceID(err)
	if errorTraceID != nil {
		fields["error.trace.id"] = errorTraceID
	}

	errorTransactionID := customError.TransactionID(err)
	if errorTransactionID != nil {
		fields["error.transaction.id"] = errorTransactionID
	}

	errorSpanID := customError.SpanID(err)
	if errorSpanID != nil {
		fields["error.span.id"] = errorSpanID
	}

	externalError := customError.ExternalError(err)
	if externalError != nil {
		fields["error.original"] = externalError
	}

	return fields
}

func getContextFields(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{}
	if ctx == nil {
		return fields
	}

	fields = apm.TraceContext(ctx)

	if fields == nil {
		fields = logrus.Fields{}
	}

	funcName := apm.GetFuncNameFromCtx(ctx)
	if funcName != nil && *funcName != "" {
		fields["func"] = *funcName
	}

	return fields
}

func getFields(object interface{}) logrus.Fields {
	fields := logrus.Fields{}

	var objectSerialized string
	if object != nil {
		switch v := object.(type) {
		case *string:
			objectSerialized = *v
		case string:
			objectSerialized = v
		case []byte:
			objectSerialized = string(v)
		default:
			if object, er := json.Marshal(object); er == nil {
				objectSerialized = string(object)
			}
		}
	}

	if objectSerialized != "" {
		fields["object"] = objectSerialized
	}

	if environment != "" {
		fields["environment"] = environment
	}

	return fields
}
