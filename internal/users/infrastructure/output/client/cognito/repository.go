package cognito

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	authPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
	cognito "github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito"
)

type (
	Config interface {
		PoolID() string
		AppClientID() string
		AppClientSecret() string
	}

	cognitoRepository struct {
		poolID          string
		appClientID     string
		appClientSecret string
		cognito         cognito.Cognito
	}
)

func NewCognitoRepository(cognito cognito.Cognito, cf Config) authPort.CognitoClientAuthRepository {
	return &cognitoRepository{
		poolID:          cf.PoolID(),
		appClientID:     cf.AppClientID(),
		appClientSecret: cf.AppClientSecret(),
		cognito:         cognito,
	}
}

func (c *cognitoRepository) getSecretHash(username string) *string {
	mac := hmac.New(sha256.New, []byte(c.appClientSecret))
	mac.Write([]byte(username + c.appClientID))
	secret := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return &secret
}
