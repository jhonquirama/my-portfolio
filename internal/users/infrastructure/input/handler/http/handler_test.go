package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ioModel "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/input/handler/http/iomodel"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	"github.com/jhonquirama/my-portfolio/pkg/mocks/services"
	ginMiddleware "github.com/jhonquirama/my-portfolio/pkg/server/gin/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewUsersHandler_UsersConfirmSignUp(t *testing.T) {
	var (
		routeURL = "/my-portfolio/users/confirm-sign-up"
		reqURL   = "/my-portfolio/users/confirm-sign-up"

		ctx     = context.TODO()
		mockCtx = mock.Anything
	)
	type (
		input struct {
			ctx                   string // se usa el context para mock
			reqURL                string
			svcUsersConfirmSignUp ioModel.UsersConfirmSignUpInput
		}
		output struct {
			svcUsersConfirmSignUpStatusCode int
			svcUsersConfirmSignUpErr        error
			body                            any
		}
		test struct {
			testName string
			in       input
			out      output
		}
	)
	tests := []test{
		{
			testName: "BAD REQUEST",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
				svcUsersConfirmSignUp: ioModel.UsersConfirmSignUpInput{
					Email: "",
					Code:  "",
				},
			},
			out: output{
				svcUsersConfirmSignUpStatusCode: 400,
				svcUsersConfirmSignUpErr:        customError.New(ctx, customError.RequestBodyValidation),
				body:                            customError.New(ctx, customError.RequestBodyValidation),
			},
		},
		{
			testName: "CONFIRM SIGNUP FAILED BY SVC",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
				svcUsersConfirmSignUp: ioModel.UsersConfirmSignUpInput{
					Email: "cccccc@gmail.com",
					Code:  "123456",
				},
			},
			out: output{
				svcUsersConfirmSignUpStatusCode: 500,
				svcUsersConfirmSignUpErr:        customError.New(ctx, customError.UnknownError),
				body:                            customError.New(ctx, customError.UnknownError),
			},
		},
		{
			testName: "CONFIRM SIGNUP SUCCESS",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
				svcUsersConfirmSignUp: ioModel.UsersConfirmSignUpInput{
					Email: "cccccc@gmail.com",
					Code:  "123456",
				},
			},
			out: output{
				svcUsersConfirmSignUpStatusCode: 200,
				svcUsersConfirmSignUpErr:        nil,
				body:                            nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()

			requestContent, _ := json.Marshal(tt.in.svcUsersConfirmSignUp)
			req, err := http.NewRequest(http.MethodPost, tt.in.reqURL, bytes.NewReader(requestContent))
			if err != nil {
				require.NoError(t, err)
			}

			svc := &mocks.UsersService{}
			handler := NewUsersHandler(svc)

			svc.On("UsersConfirmSignUp", tt.in.ctx,
				ioModel.MapUsersConfirmSignUpIOModelToSignUpModel(tt.in.svcUsersConfirmSignUp)).
				Return(tt.out.svcUsersConfirmSignUpErr)

			route := gin.Default()
			route.Use(func(c *gin.Context) {
				ginMiddleware.MiddlewareError(c)
			})
			route.POST(routeURL, handler.UsersConfirmSignUp)
			route.ServeHTTP(w, req)
			statusCode := w.Result().StatusCode
			bodyStr := w.Body.String()

			assert.Equal(t, tt.out.svcUsersConfirmSignUpStatusCode, statusCode)

			if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {
				outResError := tt.out.body.(error)
				httpResError := unmarshalResErr(bodyStr)

				assert.Equal(t, customError.Code(outResError), httpResError.Code)
			}
		})
	}
}

func TestUsersHandler_UsersSignUp(t *testing.T) {
	var (
		routeURL = "/my-portfolio/users/sign-up"
		reqURL   = "/my-portfolio/users/sign-up"

		ctx     = context.TODO()
		mockCtx = mock.Anything
	)
	type (
		input struct {
			ctx            string // se usa el context para mock
			reqURL         string
			svcUsersSignUp ioModel.UsersSignUpInput
		}
		output struct {
			svcUsersSignUpStatusCode int
			svcUsersSignUpErr        error
			body                     any
		}
		test struct {
			testName string
			in       input
			out      output
		}
	)
	tests := []test{
		{
			testName: "BAD REQUEST",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
				svcUsersSignUp: ioModel.UsersSignUpInput{
					Email:    "",
					Password: "",
				},
			},
			out: output{
				svcUsersSignUpStatusCode: 400,
				svcUsersSignUpErr:        customError.New(ctx, customError.RequestBodyValidation),
				body:                     customError.New(ctx, customError.RequestBodyValidation),
			},
		},
		{
			testName: "SIGNUP PASS WRONG",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
				svcUsersSignUp: ioModel.UsersSignUpInput{
					Email:    "cccccc@gmail.com",
					Password: "123456",
				},
			},
			out: output{
				svcUsersSignUpStatusCode: 400,
				svcUsersSignUpErr:        customError.New(ctx, customError.RequestBodyValidation),
				body:                     customError.New(ctx, customError.RequestBodyValidation),
			},
		},
		{
			testName: "SIGNUP PASS GOOD WRONG BY SVC",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
				svcUsersSignUp: ioModel.UsersSignUpInput{
					Email:    "cccccc@gmail.com",
					Password: "12345678",
				},
			},
			out: output{
				svcUsersSignUpStatusCode: 400,
				svcUsersSignUpErr:        customError.New(ctx, customError.RequestBodyValidation),
				body:                     customError.New(ctx, customError.RequestBodyValidation),
			},
		},
		{
			testName: "SIGNUP SUCCESS",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
				svcUsersSignUp: ioModel.UsersSignUpInput{
					Email:    "cccccc@gmail.com",
					Password: "12345678",
				},
			},
			out: output{
				svcUsersSignUpStatusCode: 201,
				svcUsersSignUpErr:        nil,
				body:                     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()

			requestContent, _ := json.Marshal(tt.in.svcUsersSignUp)
			req, err := http.NewRequest(http.MethodPost, tt.in.reqURL, bytes.NewReader(requestContent))
			if err != nil {
				require.NoError(t, err)
			}

			svc := &mocks.UsersService{}
			handler := NewUsersHandler(svc)

			svc.On("UsersSignUp", tt.in.ctx,
				ioModel.MapUsersSignUpIOModelToSignUpModel(tt.in.svcUsersSignUp)).
				Return(tt.out.svcUsersSignUpErr)

			route := gin.Default()
			route.Use(func(c *gin.Context) {
				ginMiddleware.MiddlewareError(c)
			})
			route.POST(routeURL, handler.UsersSignUp)
			route.ServeHTTP(w, req)
			statusCode := w.Result().StatusCode
			bodyStr := w.Body.String()

			assert.Equal(t, tt.out.svcUsersSignUpStatusCode, statusCode)

			if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {
				outResError := tt.out.body.(error)
				httpResError := unmarshalResErr(bodyStr)

				assert.Equal(t, customError.Code(outResError), httpResError.Code)
			}
		})
	}
}

func unmarshalResErr(str string) customError.ErrorResponse {
	var res customError.ErrorResponse

	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(fmt.Sprintf("error unmarshalling response error http: %v", err))
	}
	return res
}
