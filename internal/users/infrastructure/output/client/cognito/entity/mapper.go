package entity

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	cognitoIdentity "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

func SignUpSvcToCognito(data usersModel.UsersSignUpInput, clientID string) *cognitoIdentity.SignUpInput {
	return &cognitoIdentity.SignUpInput{
		ClientId:   aws.String(clientID),
		Username:   aws.String(data.UserEmail),
		Password:   aws.String(data.UserPassword),
		SecretHash: data.SecretHash,
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(data.UserEmail)},
		},
	}
}

func SignUpCognitoToSvc(data *cognitoIdentity.SignUpOutput) usersModel.UsersSignUpOutput {
	sub := ""
	if data.Session != nil {
		sub = *data.UserSub
	}

	session := ""
	if data.Session != nil {
		session = *data.Session
	}

	return usersModel.UsersSignUpOutput{
		UserConfirmed:       data.UserConfirmed,
		UserSub:             sub,
		CodeDeliveryDetails: data.CodeDeliveryDetails,
		Session:             session,
		ResultMetadata:      data.ResultMetadata,
	}
}

func ConfirmSignUpSvcToCognito(data usersModel.UsersConfirmSignUpInput,
	clientID string) *cognitoIdentity.ConfirmSignUpInput {
	return &cognitoIdentity.ConfirmSignUpInput{
		ClientId:         aws.String(clientID),
		Username:         aws.String(data.UserEmail),
		ConfirmationCode: aws.String(data.UserCode),
		SecretHash:       data.SecretHash,
	}
}

func ConfirmSignUpCognitoToSvc(data *cognitoIdentity.ConfirmSignUpOutput) usersModel.UsersConfirmSignUpOutput {
	session := ""
	if data.Session != nil {
		session = *data.Session
	}
	return usersModel.UsersConfirmSignUpOutput{
		UserSession: session,
	}
}
