package cognito

import (
	"context"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/jhonquirama/my-portfolio/internal/users/business/model"
	"github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito"
	"github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito/mocks"
	model2 "github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito/model"
	"github.com/stretchr/testify/mock"
	"reflect"
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
		}
		input struct {
			ctx  context.Context
			data model.UsersConfirmSignUpInput
		}
		output struct {
			want        model.UsersConfirmSignUpOutput
			mockCognito model2.ConfirmSignUpOutput
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
				data: model.UsersConfirmSignUpInput{
					UserEmail: "test",
					UserCode:  "122333",
				},
			},
			out: output{
				want:    model.UsersConfirmSignUpOutput{},
				wantErr: errors.New("session otp not valid"),
			},
			on: func(dep *depFields, in input, out output) {
				dep.cognito.On("ConfirmSignUp",
					in.ctx, mock.Anything).
					Return(out.mockCognito, out.wantErr)
			},
		},
		{
			name: "Case error otp not valid",
			in: input{
				ctx: ctx,
				data: model.UsersConfirmSignUpInput{
					UserEmail: "test",
					UserCode:  "122333",
				},
			},
			out: output{
				want:    model.UsersConfirmSignUpOutput{},
				wantErr: errors.New("session otp not valid"),
			},
			on: func(dep *depFields, in input, out output) {
				dep.cognito.On("ConfirmSignUp",
					in.ctx, mock.Anything).
					Return(out.mockCognito, out.wantErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &depFields{
				&mocks.Cognito{},
			}

			repo := NewCognitoRepository(f.cognito, config)

			if tt.on != nil {
				tt.on(f, tt.in, tt.out)
			}

			got, err := repo.ConfirmSignUp(tt.in.ctx, tt.in.data)

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
			if diff := cmp.Diff(tt.out.want, got); diff != "" {
				t.Errorf("❌ Diferencia en ConfirmSignUp():\n%s", diff)

				// Imprimir estructuras de manera más clara
				t.Logf("🟢 Esperado:\n%s", spew.Sdump(tt.out.want))
				t.Logf("🔴 Obtenido:\n%s", spew.Sdump(got))
			}
		})
	}
}

func Test_cognitoRepository_SignUp(t *testing.T) {
	type fields struct {
		poolID          string
		appClientID     string
		appClientSecret string
		cognito         cognito.Cognito
	}
	type args struct {
		ctx  context.Context
		data model.UsersSignUpInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.UsersSignUpOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cognitoRepository{
				poolID:          tt.fields.poolID,
				appClientID:     tt.fields.appClientID,
				appClientSecret: tt.fields.appClientSecret,
				cognito:         tt.fields.cognito,
			}
			got, err := c.SignUp(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
