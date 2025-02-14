package server

import (
	"context"
	"github.com/jhonquirama/my-portfolio/pkg/container"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	"github.com/jhonquirama/my-portfolio/pkg/monitor/observability"
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
	cnf, err := settings.NewConfigFromFile()
	if err != nil {
		if cnf, err = settings.NewConfigFromSystemManager(ctx); err != nil {
			return nil, err
		}
	}

	cnt, err := container.NewContainer(ctx, cnf)
	if err != nil {
		return nil, err
	}

	tracer, err := observability.NewProvider(ctx, cnf.ApmConfig())
	if err != nil {
		return nil, err
	}

	customLogger.SetEnvironment(cnf.ApmConfig().ApmEnvironment())

	ginServer := gin.NewGinServer(tracer, cnt)

	return &Server{
		GinServer: ginServer,
		GinLambda: lambda.NewGinLambdaServer(ginServer),
		Config:    cnf,
	}, nil
}
