package cognito

import (
	"context"
	"strings"

	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
	cognitoEntity "github.com/jhonquirama/my-portfolio/internal/users/infrastructure/output/client/cognito/entity"
	customError "github.com/jhonquirama/my-portfolio/pkg/error"
	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
)

func (c *cognitoRepository) SignUp(ctx context.Context,
	data usersModel.UsersSignUpInput) (usersModel.UsersSignUpOutput, error) {
	span, ctx := apm.NewSpan(ctx, apm.Service)
	defer span.End()

	secretHash := c.getSecretHash(data.UserName)
	data.SecretHash = secretHash

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
