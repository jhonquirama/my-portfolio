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
