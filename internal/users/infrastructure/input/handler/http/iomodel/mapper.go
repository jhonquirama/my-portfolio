package iomodel

import (
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

func MapUsersSingUpIOModelToSingUpModel(info UsersSingUpInput) usersModel.UsersSingUpInput {
	return usersModel.UsersSingUpInput{
		UserName:     info.Name,
		UserEmail:    info.Email,
		UserPassword: info.Password,
	}
}
