package server

import (
	"context"

	"github.com/jhonquirama/my-portfolio/pkg/container"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
	"github.com/jhonquirama/my-portfolio/pkg/server/gin"
	lambda "github.com/jhonquirama/my-portfolio/pkg/server/gin-lambda"
	"github.com/jhonquirama/my-portfolio/pkg/settings"
)

type Server struct {
	GinServer *gin.Server
	GinLambda *lambda.GinLambdaServer
	Config    settings.Config
}

func NewServer(ctx context.Context) (*Server, error) {
	var (
		cnf    settings.Config
		cnt    container.Container
		tracer *apm.Tracer
		err    error
	)

	if cnf, err = settings.NewConfigFromFile(); err != nil {
		if cnf, err = settings.NewConfigFromSystemManager(ctx); err != nil {
			return nil, err
		}
	}

	if cnt, err = container.NewContainer(ctx, cnf); err != nil {
		return nil, err
	}

	if tracer, err = apm.NewAPM(cnf.ApmConfig()); err != nil {
		return nil, err
	}

	tracer.Flush(nil)

	customLogger.SetEnvironment(cnf.ApmConfig().ApmEnvironment())

	ginServer := gin.NewGinServer(tracer, cnt)

	return &Server{
		GinServer: ginServer,
		GinLambda: lambda.NewGinLambdaServer(ginServer),
		Config:    cnf,
	}, nil
}
