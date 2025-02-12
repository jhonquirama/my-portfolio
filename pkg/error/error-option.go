package errors

import "strings"

type (
	Option func(r *option)

	option struct {
		skip int
		Error
	}
)

func WithMessage(messages ...string) Option {
	return func(e *option) {
		e.message = strings.Join(messages, ",")
	}
}

func WithError(err error) Option {
	return func(e *option) {
		e.externalError = err
	}
}

func WithSkip(skip int) Option {
	return func(e *option) {
		e.skip = skip
	}
}

func withHTTPCode(httpCode int) Option {
	return func(e *option) {
		e.httpCode = httpCode
	}
}

func WithSend(send bool) Option {
	return func(e *option) {
		e.send = &send
	}
}
