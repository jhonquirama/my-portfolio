package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jhonquirama/my-portfolio/internal/portfolio/business/model"
	"github.com/jhonquirama/my-portfolio/internal/portfolio/business/port/mocks"
	"github.com/jhonquirama/my-portfolio/internal/portfolio/infrastructure/input/handler/http/iomodel"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	ginMiddleware "github.com/jhonquirama/my-portfolio/pkg/server/gin/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGetHealth_GetHealth(t *testing.T) {
	var (
		routeURL = "/my-portfolio/health"
		reqURL   = "/my-portfolio/health"

		ctx     = context.TODO()
		mockCtx = mock.Anything
	)
	type (
		input struct {
			ctx    string // se usa el context para mock
			reqURL string
		}
		output struct {
			svcUsersHealthStatusCode int
			svcUsersHealthErr        error
			body                     any
			response                 model.Health
		}
		test struct {
			testName string
			in       input
			out      output
		}
	)
	tests := []test{

		{
			testName: "HEALTH FAILED BY SVC",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
			},
			out: output{
				svcUsersHealthStatusCode: 500,
				svcUsersHealthErr:        customError.New(ctx, customError.UnknownError),
				body:                     customError.New(ctx, customError.UnknownError),
				response:                 model.Health{},
			},
		},
		{
			testName: "HEALTH SUCCESS",
			in: input{
				ctx:    mockCtx,
				reqURL: reqURL,
			},
			out: output{
				svcUsersHealthStatusCode: 200,
				svcUsersHealthErr:        nil,
				body:                     nil,
				response: model.Health{
					Status: "OK",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, tt.in.reqURL, nil)
			if err != nil {
				require.NoError(t, err)
			}

			svc := &mocks.PortfolioService{}
			handler := NewPortfolioHandler(svc)

			svc.On("GetHealth",
				tt.in.ctx, iomodel.ToGetHealthModel()).Return(tt.out.response, tt.out.svcUsersHealthErr)

			route := gin.Default()
			route.Use(func(c *gin.Context) {
				ginMiddleware.MiddlewareError(c)
			})
			route.GET(routeURL, handler.GetHealth)
			route.ServeHTTP(w, req)
			statusCode := w.Result().StatusCode
			bodyStr := w.Body.String()

			assert.Equal(t, tt.out.svcUsersHealthStatusCode, statusCode)

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
