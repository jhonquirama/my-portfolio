package config

type (
	Config interface {
		ApmConfig() ApmConfig
		ServerHTTPAddress() string
		DynamodbConfig() DynamodbConfig
		DynamodbClientConfig() DynamodbClientConfig
		AwsConfig() AwsConfig
		CognitoConfig() CognitoConfig
		PortfolioService() ServiceConfig
	}

	config struct {
		Service service `json:"service" validate:"required"`
	}
)

func (cnf *config) AwsConfig() AwsConfig {
	return &cnf.Service.Aws
}
func (cnf *config) DynamodbConfig() DynamodbConfig {
	return &cnf.Service.Aws.Dynamodb
}

func (cnf *config) CognitoConfig() CognitoConfig {
	return &cnf.Service.Aws.Cognito
}

func (cnf *config) ServerHTTPAddress() string {
	return cnf.Service.Server.HTTP.Address
}

func (cnf *config) DynamodbClientConfig() DynamodbClientConfig {
	return &cnf.Service.Aws.DynamodbClient
}

func (cnf *config) ApmConfig() ApmConfig {
	return &cnf.Service.Apm
}

func (cnf *config) PortfolioService() ServiceConfig {
	return &cnf.Service.Services.Portfolio
}
