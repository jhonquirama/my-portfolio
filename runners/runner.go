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
	Run(ctx context.Context) error
}

type runner struct {
}

func NewRunner() Runner {
	runner := runner{}

	return runner.runMode()
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
