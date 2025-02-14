package iomodel

import (
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

func MapUsersSignUpIOModelToSignUpModel(info UsersSignUpInput) usersModel.UsersSignUpInput {
	return usersModel.UsersSignUpInput{
		UserName:     info.Name,
		UserEmail:    info.Email,
		UserPassword: info.Password,
	}
}
