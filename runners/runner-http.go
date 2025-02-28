package runners

import (
	"context"

	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	"github.com/jhonquirama/my-portfolio/pkg/server"
)

type httpRunner struct {
}

func (runner *httpRunner) Run(ctx context.Context) error {
	customLogger.Info(ctx, "Running as a HTTP Server ... ")

	newServer, err := server.NewServer(ctx)
	if err != nil {
		return err
	}

	return newServer.GinServer.Run(newServer.Config.ServerHTTPAddress())
}
