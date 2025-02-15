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

	AuthOTPActiveRequested             CustomCode = "AUTH_OTP_ACTIVE_REQUESTED"
	AuthOTPErrorRequest                CustomCode = "AUTH_OTP_ERROR_REQUESTED"
	AuthOTPErrorOnGenerate             CustomCode = "AUTH_OTP_ERROR_GENERATE"
	AuthOTPGenerateChallenge           CustomCode = "AUTH_OTP_ERROR_CHALLENGE"
	AuthOTPUserAlreadyExists           CustomCode = "AUTH_USER_EXIST"
	AuthOTPCognitoError                CustomCode = "AUTH_COGNITO_ERROR"
	AuthOTPError                       CustomCode = "AUTH_ERROR"
	AuthOTPSessionNotValid             CustomCode = "AUTH_SESSION_NOT_VALID"
	AuthOTPErrorDeleteItem             CustomCode = "AUTH_ERROR_DELETED_ITEM"
	AuthOTPErrorGetItem                CustomCode = "AUTH_ERROR_GET_ITEM"
	AuthOTPUserDisabled                CustomCode = "AUTH_USER_DISABLED"
	AuthOTPUserDontExist               CustomCode = "AUTH_USER_DONT_EXIST"
	AuthOTPSingUpUser                  CustomCode = "AUTH_SING_UP_USER"
	AuthOTPSingInUser                  CustomCode = "AUTH_SING_IN_USER"
	AuthOTPRFError                     CustomCode = "AUTH_TOKEN_ERROR"
	AuthOTPRFInvalid                   CustomCode = "AUTH_TOKEN_INVALID"
	AuthOTPValidationMailInvalid       CustomCode = "AUTH_MAIL_VALIDATION"
	AuthOTPPhoneFormatInvalid          CustomCode = "AUTH_PHONE_FORMAT_INVALID"
	AuthOTPUserNotFound                CustomCode = "AUTH_USER_NOT_FOUND"
	AuthOTPActiveRequestedAnotherUser  CustomCode = "AUTH_OTP_ACTIVE_REQUESTED_OTHER_USER"
	AuthOTPMailChallengeNot            CustomCode = "AUTH_OTP_CHALLENGE_ANOTHER_USER"
	AuthOTPChallengeNotMatched         CustomCode = "AUTH_OTP_CHALLENGE_NOT_MATCHED"
	AuthOTPUpdateMailFromUser          CustomCode = "AUTH_UPDATE_MAIL_CONFIRMATION"
	AuthOTPNoTokenSessionActiveMail    CustomCode = "AUTH_NOT_SESSION_ACTIVE" // nolint: gosec
	AuthOTPEmailAlreadyRegister        CustomCode = "AUTH_EMAIL_ALREADY_REGISTERED"
	AuthOTPUserEmailAlreadyRegister    CustomCode = "AUTH_USER_EMAIL_ALREADY_REGISTERED"
	AuthOTPSearchErrorCognito          CustomCode = "AUTH_SEARCH_COGNITO"
	AuthOTPUserNotFoundOnSearchUser    CustomCode = "AUTH_USER_NOT_FOUND_SEARCH"
	AuthOTPUserExist                   CustomCode = "AUTH_USER_EXIST_CREATE"
	AuthOTPSessionSaveLastLogin        CustomCode = "AUTH_SESSION_SESSION_ERROR"
	AuthOTPRefreshTokenSessionNotValid CustomCode = "AUTH_SESSION_TOKEN_NOT_VALID" // nolint:gosec
	AuthOTPMailErrorUserNotfound       CustomCode = "AUTH_OTP_MAIL_USER_NOT_FOUND_DB"
	AuthOTPCognitoDeleteUser           CustomCode = "AUTH_USER_DELETE_ERROR"
	AuthOTPCreateUserDB                CustomCode = "AUTH_ERROR_ON_CREATE_USER_POLL"
	AuthQueryProcess                   CustomCode = "AUTH_QUERY_PROCESS"
	AuthOTPCodeExpired                 CustomCode = "AUTH_CODE_EXPIRED"
	AuthOTPErrorLogin                  CustomCode = "AUTH_ERROR_LOGIN"
	AuthOTPDeleteOTPTemp               CustomCode = "AUTH_DELETE_USER_ERROR"
	AuthOTPPhoneUnsupported            CustomCode = "AUTH_UNSUPPORTED"
	AuthOTPNotifierAvailability        CustomCode = "AUTH_ERROR_NOTIFIER_NOT_AVAILABLE"
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
	AuthOTPActiveRequested: {
		message:  "user have otp sesion code active",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPErrorRequest: {
		message:  "error cognito create users",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPErrorOnGenerate: {
		message:  "internal error on OTP",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPGenerateChallenge: {
		message:  "internal error on OTP challenge",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPError: {
		message:  "error general en auth",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPSessionNotValid: {
		message:  "session otp not valid",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPErrorDeleteItem: {
		message:  "error internal dynamo",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPUserDisabled: {
		message:  "error user is disabled",
		httpCode: http.StatusUnauthorized,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPUserDontExist: {
		message:  "error user no exist",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPSingUpUser: {
		message:  "error on create user",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.InvalidArgument,
	},
	RequestBodyValidation: {
		message:  "error on request body auth",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPRFError: {
		message:  "error on refresh token",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPRFInvalid: {
		message:  "invalid refresh token",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPUserAlreadyExists: {
		message:  "error internal otp",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPCognitoError: {
		message:  "error internal cognito",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPErrorGetItem: {
		message:  "error internal dynamo",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPValidationMailInvalid: {
		message:  "error mail invalid format",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	RequestAuthNotFound: {
		message:  "header auth not found",
		httpCode: http.StatusUnauthorized,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPUserNotFound: {
		message:  "user not found",
		httpCode: http.StatusUnauthorized,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPActiveRequestedAnotherUser: {
		message:  "other user have active request mail validation",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPMailChallengeNot: {
		message:  "challengiue is not for this user",
		httpCode: http.StatusUnauthorized,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPChallengeNotMatched: {
		message:  "answers not mached",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPUpdateMailFromUser: {
		message:  "error on update user in cognito",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPNoTokenSessionActiveMail: {
		message:  "error not session active",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPEmailAlreadyRegister: {
		message:  "error email al ready registered",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPPhoneFormatInvalid: {
		message:  "error phone format",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	NotFound: {
		message:  "not found",
		httpCode: http.StatusNotFound,
		grpcCode: codes.NotFound,
	},
	AuthOTPUserNotFoundOnSearchUser: {
		message:  "user not found",
		httpCode: http.StatusNotFound,
		grpcCode: codes.NotFound,
	},
	AuthOTPSearchErrorCognito: {
		message:  "error general in cognito",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.Internal,
	},
	AuthOTPUserExist: {
		message:  "user already exist",
		httpCode: http.StatusConflict,
		grpcCode: codes.Internal,
	},
	AuthOTPUserEmailAlreadyRegister: {
		message:  "user have email registered",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPSessionSaveLastLogin: {
		message:  "user error to save session",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPRefreshTokenSessionNotValid: {
		message:  "user error session",
		httpCode: 498,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPMailErrorUserNotfound: {
		message:  "user not found in bd",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPCognitoDeleteUser: {
		message:  "user cant be deleted in user pool",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.InvalidArgument,
	},
	AuthOTPCreateUserDB: {
		message:  "user cant be created user pool",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.InvalidArgument,
	},
	AuthQueryProcess: {
		message:  "user not found",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.Internal,
	},
	AuthOTPCodeExpired: {
		message:  "code expired",
		httpCode: http.StatusNotFound,
		grpcCode: codes.Internal,
	},
	AuthOTPErrorLogin: {
		message:  "errror in login external",
		httpCode: http.StatusNotFound,
		grpcCode: codes.Internal,
	},
	AuthOTPSingInUser: {
		message:  "errror on opt login",
		httpCode: http.StatusNotFound,
		grpcCode: codes.Internal,
	},
	AuthOTPDeleteOTPTemp: {
		message:  "errror on opt deleted",
		httpCode: http.StatusInternalServerError,
		grpcCode: codes.Internal,
	},
	AuthOTPPhoneUnsupported: {
		message:  "error unsuported type auth",
		httpCode: http.StatusBadRequest,
		grpcCode: codes.Internal,
	},
	AuthOTPNotifierAvailability: {
		message:  "service notifier not available",
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
