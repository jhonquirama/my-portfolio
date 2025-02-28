package port

import (
	"context"

	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

//go:generate mockery --name UsersService --filename users_service_mock.go --output ../../../../pkg/mocks/services
type UsersService interface {
	UsersSignUp(ctx context.Context, data usersModel.UsersSignUpInput) error
	UsersConfirmSignUp(ctx context.Context,
		data usersModel.UsersConfirmSignUpInputAndSignInInput) (usersModel.UsersSignInOutput, error)
	UsersInitiateAuth(ctx context.Context,
		user usersModel.UsersConfirmSignUpInputAndSignInInput) (usersModel.UsersSignInOutput, error)
}
