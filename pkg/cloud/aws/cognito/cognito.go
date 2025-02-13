package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cognitoClient "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	cognitoModel "github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito/model"
)

type (
	Cognito interface {
		SignUp(ctx context.Context, input cognitoModel.SignUpInput) (cognitoModel.SignUpOutput, error)
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
	output, err := c.cognitoClient.SignUp(ctx, &cognitoClient.SignUpInput{
		ClientId: input.ClientId,
		Password: input.Password,
		Username: input.Username,
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(input.ClientMetadata["email"])},
		},
		SecretHash: input.SecretHash,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}
