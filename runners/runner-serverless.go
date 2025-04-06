package runners

import (
	"context"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	"github.com/jhonquirama/my-portfolio/pkg/server"
)

type serverlessRunner struct {
}

func (runner *serverlessRunner) Run(ctx context.Context) {
	customLogger.Info(ctx, "Running as Serverless ... ")

	newServer, err := server.NewServer(ctx)
	if err != nil {
		panic(err)
	}

	if err = newServer.GinLambda.Start(); err != nil {
		panic(err)
	}
}
