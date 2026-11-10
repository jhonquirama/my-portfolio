package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"

	awsCloud "github.com/jhonquirama/my-portfolio/pkg/cloud/aws"
)

const (
	appName             string = "my-portfolio"
	appNameEnvVarToFind string = "APP_NAME"
)

func NewConfigFromSystemManager(ctx context.Context) (Config, error) {
	var (
		cnf                    config
		serviceName            = appName
		ssmNameParameterToFind string
		ssm                    awsCloud.SSM
		paramValue             string
		err                    error
	)
	if envVarToFind := os.Getenv(appNameEnvVarToFind); envVarToFind != "" {
		serviceName = envVarToFind
	}
	ssmNameParameterToFind = fmt.Sprintf("/%s", serviceName)

	if ssm, err = awsCloud.NewSSM(ctx); err != nil {
		return nil, err
	}
	if paramValue, err = ssm.GetParameter(ctx, ssmNameParameterToFind); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(paramValue), &cnf); err != nil {
		return nil, err
	}
	if err := validator.New().Struct(cnf); err != nil {
		return nil, err
	}
	cnf.Service.Apm.ServiceName = serviceName

	return &cnf, nil
}

func NewConfigFromFile() (Config, error) {
	var (
		cnf            config
		serviceName    = appName
		configFileName = "service"
		configType     = "json"
		confPaths      []string
		err            error
	)

	if serviceNameToFind := os.Getenv(appNameEnvVarToFind); serviceNameToFind != "" {
		serviceName = serviceNameToFind
	}

	confPaths = []string{
		"/etc/" + serviceName,
		"$HOME/." + serviceName,
		"./settings",
		"./",
	}

	for _, dir := range confPaths {
		viper.AddConfigPath(dir)
	}

	viper.SetConfigName(configFileName)
	viper.SetConfigType(configType)

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&cnf); err != nil {
		return nil, err
	}

	if err = validator.New().Struct(cnf); err != nil {
		return nil, err
	}

	cnf.Service.Apm.ServiceName = serviceName

	return &cnf, nil
}
