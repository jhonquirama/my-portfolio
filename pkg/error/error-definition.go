package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	UnknownError CustomCode = "SERVICE_UNKNOWN_ERROR"
)

var definitions = map[CustomCode]Error{ // nolint:gochecknoglobals
	UnknownError: {
		message:  "unknown error",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.Internal,
	},
}

func getDefinition(code CustomCode) Error {
	if definition, ok := definitions[code]; ok {
		return definition
	}

	return definitions[UnknownError]
}
