package model

func UsersConfirmSignUpToUsersSignIn(input UsersConfirmSignUpInputAndSignInInput) UsersSignUpAndSignInInput {
	return UsersSignUpAndSignInInput{
		UserEmail:    input.UserEmail,
		UserPassword: input.UserPasswd,
	}
}
