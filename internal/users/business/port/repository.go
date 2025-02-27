package port

import (
	"context"
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

//go:generate mockery --name DBAuthRepository --filename db_repository_mock.go --output ../../../../pkg/mocks/repositories
//go:generate mockery --name CognitoClientAuthRepository --filename cognito_repository_mock.go --output ../../../../pkg/mocks/repositories
type (
	DBAuthRepository interface {
	}

	CognitoClientAuthRepository interface {
		SignUp(ctx context.Context,
			data usersModel.UsersSignUpInput) (usersModel.UsersSignUpOutput, error)
		ConfirmSignUp(ctx context.Context,
			data usersModel.UsersConfirmSignUpInput) (usersModel.UsersConfirmSignUpOutput, error)
	}
)
