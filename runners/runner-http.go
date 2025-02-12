package runners

import (
	"context"

	customLogger "github.com/jhonquirama/hexa-scaffolding-ms/pkg/log"
	"github.com/jhonquirama/hexa-scaffolding-ms/pkg/server"
)

type httpRunner struct {
}

func (runner *httpRunner) Run(ctx context.Context) error {
	customLogger.Info(ctx, "Running as a HTTP Server ... ")

	server, err := server.NewServer(ctx)
	if err != nil {
		return err
	}

	return server.GinServer.Run(server.Config.ServerHTTPAddress())
}
