package logic

import (
	"context"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/jhonquirama/my-portfolio/internal/users/business/model"
	"github.com/jhonquirama/my-portfolio/pkg/mocks/repositories"
	mocks2 "github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	trace2 "go.opentelemetry.io/otel/trace"
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
			req model.UsersSignUpAndSignInInput
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
				req: model.UsersSignUpAndSignInInput{
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
				req: model.UsersSignUpAndSignInInput{
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
			trace := mocks2.TelemetryProvider{}
			svc := NewUsersService(&mocks.DBAuthRepository{}, cognito, &trace)
			mockSpan := trace2.SpanFromContext(context.Background())
			trace.On("TraceStart", context.TODO(), apmService).Return(context.TODO(), mockSpan)
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
			req model.UsersConfirmSignUpInputAndSignInInput
		}
		output struct {
			errCognitoSignUp, errCognitoSignIn, svcErr error
			res                                        model.UsersSignInOutput
			cognitoInitAuthResMock                     model.UsersSignInOutput
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
				req: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail:  "ddfdfdfg@gmail.com",
					UserCode:   "123456",
					UserPasswd: "mi11224432234#C",
				},
			},
			out: output{
				errCognitoSignUp:       errors.New("user error"),
				svcErr:                 errors.New("user error"),
				res:                    model.UsersSignInOutput{},
				cognitoInitAuthResMock: model.UsersSignInOutput{},
			},
		},
		{
			testName: "error in Auth",
			in: input{
				ctx: mockCtx,
				req: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail: "ddfdfdfg@gmail.com",
					UserCode:  "12345556",
				},
			},
			out: output{
				errCognitoSignIn:       errors.New("need password"),
				svcErr:                 errors.New("need password"),
				res:                    model.UsersSignInOutput{},
				cognitoInitAuthResMock: model.UsersSignInOutput{},
			},
		},
		{
			testName: "OK",
			in: input{
				ctx: mockCtx,
				req: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail: "ddfdfdfg@gmail.com",
					UserCode:  "12345556",
				},
			},
			out: output{
				errCognitoSignUp: nil,
				res: model.UsersSignInOutput{
					AccessToken: "dffsdfsdfsdfsdfsdfsdfsdfsdfsd",
				},
				cognitoInitAuthResMock: model.UsersSignInOutput{
					AccessToken: "dffsdfsdfsdfsdfsdfsdfsdfsdfsd",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			cognito := &mocks.CognitoClientAuthRepository{}
			trace := mocks2.NewTelemetryProvider(t)
			svc := NewUsersService(&mocks.DBAuthRepository{}, cognito, trace)
			mockSpan := trace2.SpanFromContext(context.Background())
			trace.On("TraceStart", context.TODO(), apmService).Return(context.TODO(), mockSpan)
			cognito.On("ConfirmSignUp", tt.in.ctx, tt.in.req).Return(tt.out.errCognitoSignUp)
			cognito.On("UsersInitiateAuth", tt.in.ctx, mock.Anything).
				Return(tt.out.cognitoInitAuthResMock, tt.out.errCognitoSignIn)

			got, err := svc.UsersConfirmSignUp(ctx, tt.in.req)
			// Manejo de errores esperado vs real
			if err != nil {
				if err.Error() != tt.out.svcErr.Error() {
					t.Fatalf("❌ Error diferente:\n🟢 Esperado: %v\n🔴 Obtenido: %v", tt.out.errCognitoSignUp, err)
				}
			} else if tt.out.errCognitoSignUp != nil {
				t.Fatalf("❌ Se esperaba un error pero no ocurrió: %v", tt.out.errCognitoSignUp)
			}

			// Comparación de estructuras con mejor formato
			if diff := cmp.Diff(tt.out.res, got); diff != "" {
				t.Errorf("❌ Diferencia en ConfirmSignUp():\n%s", diff)

				// Imprimir estructuras de manera más clara
				t.Logf("🟢 Esperado:\n%s", spew.Sdump(tt.out.res))
				t.Logf("🔴 Obtenido:\n%s", spew.Sdump(got))
			}
		})
	}
}

func TestNewUsersService_UsersSignIn(t *testing.T) {
	var (
		ctx     = context.TODO()
		mockCtx = mock.Anything
	)
	type (
		input struct {
			ctx string // se usa el context para mock
			req model.UsersSignUpAndSignInInput
		}
		output struct {
			err error
			res model.UsersSignInOutput
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
				req: model.UsersSignUpAndSignInInput{
					UserEmail:    "ddfdfdfg@gmail.com",
					UserPassword: "sfdfgfdge54353ZX",
				},
			},
			out: output{
				err: errors.New("user error"),
				res: model.UsersSignInOutput{},
			},
		},
		{
			testName: "OK",
			in: input{
				ctx: mockCtx,
				req: model.UsersSignUpAndSignInInput{
					UserEmail:    "ddfdfdfg@gmail.com",
					UserPassword: "sfdfgfdge54353ZX",
				},
			},
			out: output{
				err: nil,
				res: model.UsersSignInOutput{
					AccessToken: "dffsdfsdfsdfsdfsdfsdfsdfsdfsd",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			cognito := &mocks.CognitoClientAuthRepository{}
			trace := mocks2.TelemetryProvider{}
			svc := NewUsersService(&mocks.DBAuthRepository{}, cognito, &trace)
			mockSpan := trace2.SpanFromContext(context.Background())
			trace.On("TraceStart", context.TODO(), apmService).Return(context.TODO(), mockSpan)
			cognito.On("UsersInitiateAuth", tt.in.ctx, tt.in.req).Return(tt.out.res, tt.out.err)

			got, err := svc.UsersInitiateAuth(ctx, tt.in.req)
			// Manejo de errores esperado vs real
			if err != nil {
				if err.Error() != tt.out.err.Error() {
					t.Fatalf("❌ Error diferente:\n🟢 Esperado: %v\n🔴 Obtenido: %v", tt.out.err, err)
				}
			} else if tt.out.err != nil {
				t.Fatalf("❌ Se esperaba un error pero no ocurrió: %v", tt.out.err)
			}

			// Comparación de estructuras con mejor formato
			if diff := cmp.Diff(tt.out.res, got); diff != "" {
				t.Errorf("❌ Diferencia en ConfirmSignUp():\n%s", diff)

				// Imprimir estructuras de manera más clara
				t.Logf("🟢 Esperado:\n%s", spew.Sdump(tt.out.res))
				t.Logf("🔴 Obtenido:\n%s", spew.Sdump(got))
			}
		})
	}
}
