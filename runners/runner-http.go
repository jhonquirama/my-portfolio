package runners

import (
	"context"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
)

type httpRunner struct {
}

func (runner *httpRunner) Run(ctx context.Context) {
	customLogger.Info(ctx, "Running as a HTTP Server ... ")
}
