package aws

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSM interface {
	GetParameter(ctx context.Context, parameterInput string) (string, error)
}

type ssmanager struct {
	client *ssm.Client
}

func NewSSM(ctx context.Context) (SSM, error) {
	var (
		conf aws.Config
		err  error
	)

	if conf, err = config.LoadDefaultConfig(ctx); err != nil {
		return nil, err
	}

	return &ssmanager{
		client: ssm.NewFromConfig(conf),
	}, nil
}

func (s ssmanager) GetParameter(ctx context.Context, parameterInput string) (string, error) {
	var (
		output *ssm.GetParameterOutput
		err    error
	)

	input := ssm.GetParameterInput{
		Name:           aws.String(parameterInput),
		WithDecryption: aws.Bool(false),
	}
	if output, err = s.client.GetParameter(ctx, &input); err != nil {
		return "", err
	}
	if output == nil || output.Parameter == nil || output.Parameter.Value == nil {
		return "", errors.New("aws ssm output param value is nill")
	}

	return *output.Parameter.Value, nil
}
