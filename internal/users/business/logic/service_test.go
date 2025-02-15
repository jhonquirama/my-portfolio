package logic

import (
	"context"
	"errors"
	"github.com/jhonquirama/my-portfolio/internal/users/business/model"
	"github.com/jhonquirama/my-portfolio/internal/users/business/port/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewUsersService_UsersSignUp(t *testing.T) {
	var (
		ctx     = context.TODO()
		mockCtx = mock.Anything
	)
	type (
		input struct {
			ctx string // se usa el context para mock
			req model.UsersSignUpInput
		}
		output struct {
			err error
			res model.UsersSignUpOutput
		}
		test struct {
			testName string
			in       input
			out      output
		}
	)
	tests := []test{
		{
			testName: "WRONG",
			in: input{
				ctx: mockCtx,
				req: model.UsersSignUpInput{
					UserEmail:    "ddfdfdfg@gmail.com",
					UserPassword: "sfdfgfdge54353ZX",
				},
			},
			out: output{
				err: errors.New("user error"),
				res: model.UsersSignUpOutput{},
			},
		},
		{
			testName: "OK",
			in: input{
				ctx: mockCtx,
				req: model.UsersSignUpInput{
					UserEmail:    "ddfdfdfg@gmail.com",
					UserPassword: "sfdfgfdge54353ZX",
				},
			},
			out: output{
				err: nil,
				res: model.UsersSignUpOutput{
					Session: "dffsdfsdfsdfsdfsdfsdfsdfsdfsd",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			cognito := &mocks.CognitoClientAuthRepository{}
			svc := NewUsersService(&mocks.DBAuthRepository{}, cognito)

			cognito.On("SignUp", tt.in.ctx, tt.in.req).Return(tt.out.res, tt.out.err)

			err := svc.UsersSignUp(ctx, tt.in.req)
			assert.Equal(t, err, tt.out.err)
		})
	}
}

func TestNewUsersService_UsersConfirmSignUp(t *testing.T) {
	var (
		ctx     = context.TODO()
		mockCtx = mock.Anything
	)
	type (
		input struct {
			ctx string // se usa el context para mock
			req model.UsersConfirmSignUpInput
		}
		output struct {
			err error
			res model.UsersConfirmSignUpOutput
		}
		test struct {
			testName string
			in       input
			out      output
		}
	)
	tests := []test{
		{
			testName: "WRONG",
			in: input{
				ctx: mockCtx,
				req: model.UsersConfirmSignUpInput{
					UserEmail: "ddfdfdfg@gmail.com",
					UserCode:  "123456",
				},
			},
			out: output{
				err: errors.New("user error"),
				res: model.UsersConfirmSignUpOutput{},
			},
		},
		{
			testName: "OK",
			in: input{
				ctx: mockCtx,
				req: model.UsersConfirmSignUpInput{
					UserEmail: "ddfdfdfg@gmail.com",
					UserCode:  "12345556",
				},
			},
			out: output{
				err: nil,
				res: model.UsersConfirmSignUpOutput{
					UserSession: "dffsdfsdfsdfsdfsdfsdfsdfsdfsd",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			cognito := &mocks.CognitoClientAuthRepository{}
			svc := NewUsersService(&mocks.DBAuthRepository{}, cognito)

			cognito.On("ConfirmSignUp", tt.in.ctx, tt.in.req).Return(tt.out.res, tt.out.err)

			err := svc.UsersConfirmSignUp(ctx, tt.in.req)
			assert.Equal(t, err, tt.out.err)
		})
	}
}
