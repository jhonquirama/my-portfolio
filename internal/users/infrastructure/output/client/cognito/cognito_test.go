package cognito

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/jhonquirama/my-portfolio/internal/users/business/model"
	"github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito/mocks"
	model2 "github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito/model"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	mocks2 "github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel/mocks"
	"github.com/stretchr/testify/mock"
	trace2 "go.opentelemetry.io/otel/trace"
	"testing"
)

type (
	cognitoConfig struct{}
)

var (
	config = &cognitoConfig{}
)

func (db *cognitoConfig) PoolID() string {
	return "POOLID"
}
func (db *cognitoConfig) AppClientID() string {
	return "APPClientID"
}
func (db *cognitoConfig) AppClientSecret() string {
	return "APPClientSecret"
}

func Test_cognitoRepository_ConfirmSignUp(t *testing.T) {
	var (
		ctx = context.Background()
	)
	type (
		depFields struct {
			cognito *mocks.Cognito
			trace   *mocks2.TelemetryProvider
		}
		input struct {
			ctx  context.Context
			data model.UsersConfirmSignUpInputAndSignInInput
		}
		output struct {
			want             model.UsersSignInOutput
			mockCognito      model2.ConfirmSignUpOutput
			wantErr          error
			cognitoErrorWant error
		}
	)
	tests := []struct {
		name string
		in   input
		out  output
		on   func(dep *depFields, in input, out output)
	}{
		{
			name: "Case error otp not valid",
			in: input{
				ctx: ctx,
				data: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail: "test",
					UserCode:  "122333",
				},
			},
			out: output{
				want:    model.UsersSignInOutput{},
				wantErr: errors.New("user auth error"),
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)

				dep.cognito.On("ConfirmSignUp", mock.Anything, mock.Anything).
					Return(nil, out.wantErr)
			},
		},
		{
			name: "Case error user Exist",
			in: input{
				ctx: ctx,
				data: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail: "test",
					UserCode:  "122333",
				},
			},
			out: output{
				want:             model.UsersSignInOutput{},
				wantErr:          errors.New("user auth error"),
				cognitoErrorWant: errors.New("UsernameExistsException"),
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)

				dep.cognito.On("ConfirmSignUp", mock.Anything, mock.Anything).
					Return(out.mockCognito, out.cognitoErrorWant)
			},
		},
		{
			name: "Case Success",
			in: input{
				ctx: ctx,
				data: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail: "test",
					UserCode:  "122333",
				},
			},
			out: output{
				want:        model.UsersSignInOutput{},
				wantErr:     nil,
				mockCognito: nil,
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)
				mockOutput := model2.ConfirmSignUpOutput(&cognitoidentityprovider.ConfirmSignUpOutput{}) // 🔥 Conversión explícita

				dep.cognito.On("ConfirmSignUp", mock.Anything, mock.Anything).
					Return(mockOutput, out.wantErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &depFields{
				&mocks.Cognito{},
				&mocks2.TelemetryProvider{},
			}

			repo := NewCognitoRepository(f.trace, f.cognito, config)

			if tt.on != nil {
				tt.on(f, tt.in, tt.out)
			}

			err := repo.ConfirmSignUp(tt.in.ctx, tt.in.data)

			// Manejo de errores esperado vs real
			if err != nil {
				if tt.out.wantErr == nil {
					t.Fatalf("❌ Error inesperado: %v", err)
				}
				if err.Error() != tt.out.wantErr.Error() {
					t.Fatalf("❌ Error diferente:\n🟢 Esperado: %v\n🔴 Obtenido: %v", tt.out.wantErr, err)
				}
			} else if tt.out.wantErr != nil {
				t.Fatalf("❌ Se esperaba un error pero no ocurrió: %v", tt.out.wantErr)
			}
		})
	}
}

func Test_cognitoRepository_SignUp(t *testing.T) {
	var (
		ctx = context.Background()
	)
	type (
		depFields struct {
			cognito *mocks.Cognito
			trace   *mocks2.TelemetryProvider
		}
		input struct {
			ctx  context.Context
			data model.UsersSignUpInput
		}
		output struct {
			want             model.UsersSignUpOutput
			mockCognito      model2.SignUpOutput
			wantErr          error
			cognitoErrorWant error
		}
	)
	tests := []struct {
		name string
		in   input
		out  output
		on   func(dep *depFields, in input, out output)
	}{
		{
			name: "Case error otp not valid",
			in: input{
				ctx: ctx,
				data: model.UsersSignUpInput{
					UserEmail:    "test",
					UserPassword: "122333",
				},
			},
			out: output{
				want:    model.UsersSignUpOutput{},
				wantErr: errors.New("user auth error"),
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)

				dep.cognito.On("SignUp", mock.Anything, mock.Anything).
					Return(nil, out.wantErr)
			},
		},
		{
			name: "Case error user Exist",
			in: input{
				ctx: ctx,
				data: model.UsersSignUpInput{
					UserEmail:    "test",
					UserPassword: "122333",
				},
			},
			out: output{
				want: model.UsersSignUpOutput{},
				wantErr: customError.New(ctx, customError.AuthErrorRequested,
					customError.WithError(errors.New("error on create user"))),
				cognitoErrorWant: errors.New("UsernameExistsException"),
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)

				dep.cognito.On("SignUp", mock.Anything, mock.Anything).
					Return(out.mockCognito, out.cognitoErrorWant)
			},
		},
		{
			name: "Case Success",
			in: input{
				ctx: ctx,
				data: model.UsersSignUpInput{
					UserEmail:    "test",
					UserPassword: "122333",
				},
			},
			out: output{
				want:        model.UsersSignUpOutput{},
				wantErr:     nil,
				mockCognito: nil,
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)
				mockOutput := model2.SignUpOutput(&cognitoidentityprovider.SignUpOutput{}) // 🔥 Conversión explícita

				dep.cognito.On("SignUp", mock.Anything, mock.Anything).
					Return(mockOutput, out.wantErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &depFields{
				&mocks.Cognito{},
				&mocks2.TelemetryProvider{},
			}

			repo := NewCognitoRepository(f.trace, f.cognito, config)

			if tt.on != nil {
				tt.on(f, tt.in, tt.out)
			}

			got, err := repo.SignUp(tt.in.ctx, tt.in.data)

			// Manejo de errores esperado vs real
			if err != nil {
				if tt.out.wantErr == nil {
					t.Fatalf("❌ Error inesperado: %v", err)
				}
				if err.Error() != tt.out.wantErr.Error() {
					t.Fatalf("❌ Error diferente:\n🟢 Esperado: %v\n🔴 Obtenido: %v", tt.out.wantErr, err)
				}
			} else if tt.out.wantErr != nil {
				t.Fatalf("❌ Se esperaba un error pero no ocurrió: %v", tt.out.wantErr)
			}

			// Comparación de estructuras con mejor formato
			if diff := cmp.Diff(tt.out.want.Session, got.Session); diff != "" {
				t.Errorf("❌ Diferencia en ConfirmSignUp():\n%s", diff)

				// Imprimir estructuras de manera más clara
				t.Logf("🟢 Esperado:\n%s", spew.Sdump(tt.out.want))
				t.Logf("🔴 Obtenido:\n%s", spew.Sdump(got))
			}
		})
	}
}

func Test_cognitoRepository_UsersInitiateAuth(t *testing.T) {
	var (
		ctx = context.Background()
	)
	type (
		depFields struct {
			cognito *mocks.Cognito
			trace   *mocks2.TelemetryProvider
		}
		input struct {
			ctx  context.Context
			data model.UsersConfirmSignUpInputAndSignInInput
		}
		output struct {
			want        model.UsersSignInOutput
			mockCognito model2.InitiateAuthOutput
			wantErr     error
		}
	)
	tests := []struct {
		name string
		in   input
		out  output
		on   func(dep *depFields, in input, out output)
	}{
		{
			name: "Case error otp not valid",
			in: input{
				ctx: ctx,
				data: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail: "test",
					UserCode:  "122333",
				},
			},
			out: output{
				want:    model.UsersSignInOutput{},
				wantErr: errors.New("user auth error"),
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)

				dep.cognito.On("SignIn", mock.Anything, mock.Anything).
					Return(out.mockCognito, out.wantErr)
			},
		},
		{
			name: "Case Success",
			in: input{
				ctx: ctx,
				data: model.UsersConfirmSignUpInputAndSignInInput{
					UserEmail:  "test",
					UserCode:   "122333",
					UserPasswd: "test",
				},
			},
			out: output{
				want:        model.UsersSignInOutput{Token: "test"},
				wantErr:     nil,
				mockCognito: nil,
			},
			on: func(dep *depFields, in input, out output) {
				mockSpan := trace2.SpanFromContext(context.Background())
				dep.trace.On("TraceStart", mock.Anything, apmCognito).Return(context.TODO(), mockSpan)
				mockOutput := model2.InitiateAuthOutput(&cognitoidentityprovider.InitiateAuthOutput{
					AuthenticationResult: &types.AuthenticationResultType{
						AccessToken: aws.String("test"),
					},
				}) // 🔥 Conversión explícita

				dep.cognito.On("SignIn", mock.Anything, mock.Anything).
					Return(mockOutput, out.wantErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &depFields{
				&mocks.Cognito{},
				&mocks2.TelemetryProvider{},
			}

			repo := NewCognitoRepository(f.trace, f.cognito, config)

			if tt.on != nil {
				tt.on(f, tt.in, tt.out)
			}

			got, err := repo.UsersInitiateAuth(tt.in.ctx, tt.in.data)

			// Manejo de errores esperado vs real
			if err != nil {
				if tt.out.wantErr == nil {
					t.Fatalf("❌ Error inesperado: %v", err)
				}
				if err.Error() != tt.out.wantErr.Error() {
					t.Fatalf("❌ Error diferente:\n🟢 Esperado: %v\n🔴 Obtenido: %v", tt.out.wantErr, err)
				}
			} else if tt.out.wantErr != nil {
				t.Fatalf("❌ Se esperaba un error pero no ocurrió: %v", tt.out.wantErr)
			}

			if diff := cmp.Diff(tt.out.want.Token, got.Token); diff != "" {
				t.Errorf("❌ Diferencia en ConfirmSignUp():\n%s", diff)

				// Imprimir estructuras de manera más clara
				t.Logf("🟢 Esperado:\n%s", spew.Sdump(tt.out.want))
				t.Logf("🔴 Obtenido:\n%s", spew.Sdump(got))
			}
		})
	}
}
