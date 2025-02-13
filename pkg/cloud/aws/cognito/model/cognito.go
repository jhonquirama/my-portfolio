package model

import (
	cognitoIdentity "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type (
	SignUpInput  *cognitoIdentity.SignUpInput
	SignUpOutput *cognitoIdentity.SignUpOutput
)
