package cognito

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	authPort "github.com/jhonquirama/my-portfolio/internal/users/business/port"
	cognito "github.com/jhonquirama/my-portfolio/pkg/cloud/aws/cognito"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability/gotel"
)

const (
	apmCognito = "cognito"
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
		apm             gotel.TelemetryProvider
	}
)

func NewCognitoRepository(apm gotel.TelemetryProvider,
	cognito cognito.Cognito, cf Config) authPort.CognitoClientAuthRepository {
	return &cognitoRepository{
		apm:             apm,
		poolID:          cf.PoolID(),
		appClientID:     cf.AppClientID(),
		appClientSecret: cf.AppClientSecret(),
		cognito:         cognito,
	}
}

func (c *cognitoRepository) getSecretHash(username string) *string {
	hash := hmac.New(sha256.New, []byte(c.appClientSecret))
	hash.Write([]byte(username + c.appClientID))
	secret := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return &secret
}
