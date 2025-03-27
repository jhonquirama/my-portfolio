package iomodel

import (
	usersModel "github.com/jhonquirama/my-portfolio/internal/users/business/model"
)

func MapUsersSignUpIOModelToSignUpModel(info UsersSignUpAndSingInInput) usersModel.UsersSignUpAndSignInInput {
	return usersModel.UsersSignUpAndSignInInput{
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
		AccessToken:      info.AccessToken,
		RefreshToken:     info.RefreshToken,
		TokenID:          info.TokenID,
		ExpiresInSeconds: info.ExpiresInSeconds,
	}
}
