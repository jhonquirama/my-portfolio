package cognito

import (
	"context"
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
	cognitoEntity "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/output/client/cognito/entity"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
)

func (c *cognitoRepository) SignUp(ctx context.Context,
	data usersModel.UsersSignUpInput) (usersModel.UsersSignUpOutput, error) {
	ctx, span := c.apm.TraceStart(ctx, apmCognito)
	defer span.End()

	data.SecretHash = c.getSecretHash(data.UserEmail)

	signUpResult, err := c.cognito.SignUp(ctx, cognitoEntity.SignUpSvcToCognito(data, c.appClientID))
	if err != nil {
		return usersModel.UsersSignUpOutput{},
			customError.New(ctx, customError.AuthErrorRequested, customError.WithError(err))
	}

	return cognitoEntity.SignUpCognitoToSvc(signUpResult), nil
}

func (c *cognitoRepository) ConfirmSignUp(ctx context.Context, data usersModel.UsersConfirmSignUpInputAndSignInInput,
) error {
	ctx, span := c.apm.TraceStart(ctx, apmCognito)
	defer span.End()

	data.SecretHash = c.getSecretHash(data.UserEmail)

	_, err := c.cognito.
		ConfirmSignUp(ctx, cognitoEntity.ConfirmSignUpSvcToCognito(data, c.appClientID))
	if err != nil {
		return customError.New(ctx, customError.AuthErrorRequested, customError.WithError(err))
	}

	return nil
}

func (c *cognitoRepository) UsersInitiateAuth(ctx context.Context,
	data usersModel.UsersConfirmSignUpInputAndSignInInput) (usersModel.UsersSignInOutput, error) {
	ctx, span := c.apm.TraceStart(ctx, apmCognito)
	defer span.End()

	data.SecretHash = c.getSecretHash(data.UserEmail)

	initAuthResult, err := c.cognito.SignIn(ctx, cognitoEntity.InitiateAuthSvcToCognito(data, c.appClientID))
	if err != nil {
		return usersModel.UsersSignInOutput{},
			customError.New(ctx, customError.AuthErrorRequested, customError.WithError(err))
	}

	return cognitoEntity.InitiateAuthCognitoResToSvc(initAuthResult), nil
}
