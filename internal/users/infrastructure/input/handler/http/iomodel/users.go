package iomodel

type (
	UsersSignUpInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=256"`
	}

	UsersConfirmSignUpInput struct {
		Email  string `json:"email" binding:"required,email"`
		Code   string `json:"code" binding:"required"`
		Passwd string `json:"password" binding:"required,min=8,max=256"`
	}

	UsersSignInOutput struct {
		Token string `json:"token"`
	}
)
