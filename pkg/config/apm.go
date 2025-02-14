package config // nolint: dupl

type (
	apm struct {
		ServiceName string `mapstructure:"service_name" json:"service_name"`
		ServerHost  string `mapstructure:"server_host" json:"server_host" validate:"required"`
		SecretToken string `mapstructure:"secret_token" json:"secret_token" validate:"required"`
		LogFile     string `mapstructure:"log_file" json:"log_file" validate:"required"`
		LogLevel    string `mapstructure:"log_level" json:"log_level" validate:"required"`
		Environment string `json:"environment" validate:"required"`
	}

	ApmConfig interface {
		ApmServiceName() string
		ApmServerHost() string
		ApmSecretToken() string
		ApmLogFile() string
		ApmLogLevel() string
		ApmEnvironment() string
	}
)

func (apm *apm) ApmServiceName() string {
	return apm.ServiceName
}

func (apm *apm) ApmServerHost() string {
	return apm.ServerHost
}

func (apm *apm) ApmSecretToken() string {
	return apm.SecretToken
}

func (apm *apm) ApmLogFile() string {
	return apm.LogFile
}

func (apm *apm) ApmLogLevel() string {
	return apm.LogLevel
}

func (apm *apm) ApmEnvironment() string {
	return apm.Environment
}
