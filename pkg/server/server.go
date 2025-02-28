package server

import (
	"context"
	"github.com/jhonquirama/my-portfolio/pkg/config"
	"github.com/jhonquirama/my-portfolio/pkg/container"
	customLogger "github.com/jhonquirama/my-portfolio/pkg/log"
	"github.com/jhonquirama/my-portfolio/pkg/server/gin"
	lambda "github.com/jhonquirama/my-portfolio/pkg/server/gin-lambda"
)

type Server struct {
	GinServer *gin.Server
	GinLambda *lambda.GinLambdaServer
	Config    config.Config
}

func NewServer(ctx context.Context) (*Server, error) {
	cnf, err := config.NewConfigFromFile()
	if err != nil {
		if cnf, err = config.NewConfigFromSystemManager(ctx); err != nil {
			return nil, err
		}
	}

	cnt, err := container.NewContainer(ctx, cnf)
	if err != nil {
		return nil, err
	}

	customLogger.SetEnvironment(cnf.ApmConfig().ApmEnvironment())

	ginServer := gin.NewGinServer(cnt.Tracer(), cnt)

	return &Server{
		GinServer: ginServer,
		GinLambda: lambda.NewGinLambdaServer(ginServer),
		Config:    cnf,
	}, nil
}
