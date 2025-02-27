package cognito

import (
	"context"
	"strings"

	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
	cognitoEntity "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/output/client/cognito/entity"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability"
)

func (c *cognitoRepository) SignUp(ctx context.Context,
	data usersModel.UsersSignUpInput) (usersModel.UsersSignUpOutput, error) {
	ctx, span := observability.NewSpan(ctx, observability.Client)
	defer span.End()

	data.SecretHash = c.getSecretHash(data.UserEmail)

	signUpResult, err := c.cognito.SignUp(ctx, cognitoEntity.SignUpSvcToCognito(data, c.appClientID))
	if err != nil {
		if strings.Contains(err.Error(), "UsernameExistsException") {
			return usersModel.UsersSignUpOutput{},
				customError.New(ctx, customError.AuthOTPUserExist, customError.WithError(err))
		}
		return usersModel.UsersSignUpOutput{},
			customError.New(ctx, customError.AuthOTPSingUpUser, customError.WithError(err))
	}

	return cognitoEntity.SignUpCognitoToSvc(signUpResult), nil
}

func (c *cognitoRepository) ConfirmSignUp(ctx context.Context, data usersModel.UsersConfirmSignUpInput,
) (usersModel.UsersConfirmSignUpOutput, error) {
	ctx, span := observability.NewSpan(ctx, observability.Client)
	defer span.End()

	data.SecretHash = c.getSecretHash(data.UserEmail)

	confirmSignUpResult, err := c.cognito.
		ConfirmSignUp(ctx, cognitoEntity.ConfirmSignUpSvcToCognito(data, c.appClientID))
	if err != nil {
		if strings.Contains(err.Error(), "UsernameExistsException") {
			return usersModel.UsersConfirmSignUpOutput{},
				customError.New(ctx, customError.AuthOTPUserExist, customError.WithError(err))
		}
		return usersModel.UsersConfirmSignUpOutput{},
			customError.New(ctx, customError.AuthOTPSessionNotValid, customError.WithError(err))
	}

	return cognitoEntity.ConfirmSignUpCognitoToSvc(confirmSignUpResult), nil
}
