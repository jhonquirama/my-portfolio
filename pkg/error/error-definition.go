package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	UnknownError CustomCode = "SERVICE_UNKNOWN_ERROR"
	NotFound     CustomCode = "NOT_FOUND"
	QueryProcess CustomCode = "QUERY_PROCESS"

	RequestAuthNotFound   CustomCode = "REQUEST_HEADER_AUTH_VALIDATION"
	RequestBodyValidation CustomCode = "REQUEST_BODY_VALIDATION"

	AuthErrorRequested CustomCode = "AUTH_ERROR_REQUESTED"
)

var definitions = map[CustomCode]Error{ // nolint:gochecknoglobals
	UnknownError: {
		message:  "unknown error",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.Internal,
	},
	QueryProcess: {
		message:  "query error",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.Internal,
	},

	AuthErrorRequested: {
		message:  "user auth error",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},

	RequestAuthNotFound: {
		message:  "header auth not found",
		httpCode: http.StatusUnauthorized,
		grpcCode: codes.InvalidArgument,
	},

	NotFound: {
		message:  "not found",
		httpCode: http.StatusNotFound,
		grpcCode: codes.NotFound,
	},
}

func getDefinition(code CustomCode) Error {
	if definition, ok := definitions[code]; ok {
		return definition
	}

	return definitions[UnknownError]
}
