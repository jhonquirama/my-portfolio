package iomodel

import (
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

func MapUsersSignUpIOModelToSignUpModel(info UsersSignUpInput) usersModel.UsersSignUpInput {
	return usersModel.UsersSignUpInput{
		UserEmail:    info.Email,
		UserPassword: info.Password,
	}
}

func MapUsersConfirmSignUpIOModelToSignUpModel(info UsersConfirmSignUpInput,
) usersModel.UsersConfirmSignUpInputAndSignInInput {
	return usersModel.UsersConfirmSignUpInputAndSignInInput{
		UserEmail:  info.Email,
		UserCode:   info.Code,
		UserPasswd: info.Passwd,
	}
}

func MapUsersSignInModelToSignInIOModel(info usersModel.UsersSignInOutput) UsersSignInOutput {
	return UsersSignInOutput{
		Token: info.Token,
	}
}
