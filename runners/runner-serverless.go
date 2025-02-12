package runners

import (
	"context"

	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	"github.com/jhonquirama/my-portfolio/pkg/server"
)

type serverlessRunner struct {
}

func (runner *serverlessRunner) Run(ctx context.Context) error {
	customLogger.Info(ctx, "Running as Serverless ... ")

	server, err := server.NewServer(ctx)
	if err != nil {
		return err
	}

	return server.GinLambda.Start()
}
