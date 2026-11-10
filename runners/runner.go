package runners

import (
	"context"
	"os"
	"strings"
)

const (
	lambdaRuntimeAPIEnv = "AWS_LAMBDA_RUNTIME_API"
)

type Runner interface {
	Run(ctx context.Context)
}

type runner struct {
}

func NewRunner() Runner {
	newRunner := runner{}

	return newRunner.runMode()
}

func (s *runner) runMode() Runner {
	if lambdaRuntimeAPI := os.Getenv(lambdaRuntimeAPIEnv); lambdaRuntimeAPI != "" {
		return &serverlessRunner{}
	}

	switch strings.ToUpper(os.Getenv("APP_RUN_MODE")) {
	case "HTTP":
		return &httpRunner{}
	case "SERVERLESS", "LAMBDA":
		return &serverlessRunner{}
	default:
		return &defaultRunner{}
	}
}
