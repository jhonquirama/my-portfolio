package model

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
)

type (
	UsersSignUpInput struct {
		UserName       string
		UserEmail      string
		UserPassword   string
		SecretHash     *string
		UserAttributes map[string]string
	}

	UsersSignUpOutput struct {
		UserConfirmed       bool
		UserSub             string
		CodeDeliveryDetails *types.CodeDeliveryDetailsType
		Session             string
		ResultMetadata      middleware.Metadata
	}

	UsersConfirmSignUpInputAndSignInInput struct {
		UserEmail  string
		UserCode   string
		UserPasswd string
		SecretHash *string
	}

	UsersSignInOutput struct {
		Token string
	}
)
