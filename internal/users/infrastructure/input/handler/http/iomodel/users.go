package iomodel

type (
	UsersSignUpAndSingInInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=256"`
	}

	UsersConfirmSignUpInput struct {
		Email  string `json:"email" binding:"required,email"`
		Code   string `json:"code" binding:"required"`
		Passwd string `json:"password" binding:"required,min=8,max=256"`
	}

	UsersSignInOutput struct {
		AccessToken      string `json:"access_token"`
		RefreshToken     string `json:"refresh_token"`
		TokenID          string `json:"token_id"`
		ExpiresInSeconds int32  `json:"expires_in_seconds"`
	}
)
