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

func MapUsersConfirmSignUpIOModelToSignUpModel(info UsersConfirmSignUpInput) usersModel.UsersConfirmSignUpInput {
	return usersModel.UsersConfirmSignUpInput{
		UserEmail: info.Email,
		UserCode:  info.Code,
	}
}
