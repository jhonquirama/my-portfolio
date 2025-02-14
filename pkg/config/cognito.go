package config

type (
	cognito struct {
		CognitoPoolID          string `mapstructure:"pool_id" json:"pool_id" validate:"required"`
		CognitoAppClientID     string `mapstructure:"app_client_id" json:"app_client_id" validate:"required"`
		CognitoAppClientSecret string `mapstructure:"app_client_secret" json:"app_client_secret" validate:"required"`
	}
	CognitoConfig interface {
		PoolID() string
		AppClientID() string
		AppClientSecret() string
	}
)

func (c *cognito) PoolID() string {
	return c.CognitoPoolID
}

func (c *cognito) AppClientID() string {
	return c.CognitoAppClientID
}

func (c *cognito) AppClientSecret() string {
	return c.CognitoAppClientSecret
}
