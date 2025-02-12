package runners

import (
	"context"
	"fmt"

	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	"github.com/jhonquirama/my-portfolio/pkg/server"
)

type defaultRunner struct {
}

func (runner *defaultRunner) Run(ctx context.Context) error {
	server, err := server.NewServer(ctx)
	if err != nil {
		return err
	}

	// TODO: Add channel to stop gracefully

	if server.Config.ServerHTTPAddress() != "" {
		customLogger.Info(ctx, fmt.Sprintf("Running as a HTTP Server on: %s.", server.Config.ServerHTTPAddress()))

		return server.GinServer.Run(server.Config.ServerHTTPAddress())
	}

	return nil
}
