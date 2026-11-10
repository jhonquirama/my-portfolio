package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	cognitoClient "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	cognitoModel "github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito/model"
)

//go:generate mockery --name Cognito
type (
	Cognito interface {
		SignUp(ctx context.Context, input cognitoModel.SignUpInput) (cognitoModel.SignUpOutput, error)
		ConfirmSignUp(ctx context.Context, input cognitoModel.ConfirmSignUpInput) (cognitoModel.ConfirmSignUpOutput, error)
		SignIn(ctx context.Context, input cognitoModel.InitiateAuthInput,
		) (cognitoModel.InitiateAuthOutput, error)
	}
	cognito struct {
		cognitoClient *cognitoClient.Client
	}
)

func NewCognito(ctx context.Context) (Cognito, error) {
	conf, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &cognito{
		cognitoClient: cognitoClient.NewFromConfig(conf),
	}, nil
}

func (c *cognito) SignUp(ctx context.Context, input cognitoModel.SignUpInput) (cognitoModel.SignUpOutput, error) {
	return c.cognitoClient.SignUp(ctx, input)
}

func (c *cognito) ConfirmSignUp(ctx context.Context, input cognitoModel.ConfirmSignUpInput,
) (cognitoModel.ConfirmSignUpOutput, error) {
	return c.cognitoClient.ConfirmSignUp(ctx, input)
}

func (c *cognito) SignIn(ctx context.Context, input cognitoModel.InitiateAuthInput,
) (cognitoModel.InitiateAuthOutput, error) {
	return c.cognitoClient.InitiateAuth(ctx, input)
}
