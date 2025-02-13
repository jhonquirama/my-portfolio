package logic

import (
	"context"
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
	usersPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
)

type (
	Config       interface{}
	usersService struct {
		dbAuthRepository usersPort.DBAuthRepository
		usersRepository  usersPort.UsersRepository
		cognitoAuth      usersPort.CognitoClientAuthRepository
	}
)

func NewUsersService(usersRepository usersPort.UsersRepository, dbAuthRepository usersPort.DBAuthRepository,
	cognitoAuth usersPort.CognitoClientAuthRepository,
) usersPort.UsersService {
	return &usersService{
		dbAuthRepository: dbAuthRepository,
		cognitoAuth:      cognitoAuth,
		usersRepository:  usersRepository,
	}
}

func (svc *usersService) UsersSignUp(ctx context.Context,
	user usersModel.UsersSignUpInput) error {
	_, err := svc.cognitoAuth.SignUp(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
