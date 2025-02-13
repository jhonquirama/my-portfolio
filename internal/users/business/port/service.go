package port

import (
	"context"

	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

//go:generate mockery --name UsersService
type UsersService interface {
	UsersSignUp(ctx context.Context, data usersModel.UsersSignUpInput) error
}
