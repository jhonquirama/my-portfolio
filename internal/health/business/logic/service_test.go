package logic

import (
	"context"
	"github.com/jhonquirama/my-portfolio/internal/health/business/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewGetHealth_GetHealth(t *testing.T) {
	var (
		ctx     = context.TODO()
		mockCtx = mock.Anything
	)
	type (
		input struct {
			ctx string // se usa el context para mock
			req model.GetHealth
		}
		output struct {
			response model.Health
		}
		test struct {
			testName string
			in       input
			out      output
		}
	)
	tests := []test{

		{
			testName: "OK",
			in: input{
				ctx: mockCtx,
				req: model.GetHealth{},
			},
			out: output{
				response: model.Health{Status: "Ok"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			svc := NewHealthService()

			health, err := svc.GetHealth(ctx, tt.in.req)
			if err != nil {
				return
			}
			assert.NoError(t, nil, err)
			assert.Equal(t, health, tt.out.response)
		})
	}
}
