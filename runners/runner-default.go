package runners

import (
	"context"
	"github.com/jhonquirama/my-portfolio/pkg/server"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type defaultRunner struct {
}

func (runner *defaultRunner) Run(ctx context.Context) {
	newServer, err := server.NewServer(ctx)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	var quit = make(chan struct{})
	wg.Add(1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go newServer.GinServer.Run(ctx, &wg, quit, newServer.Config.ServerHTTPAddress())

	go func() {
		<-c
		log.Println("==========+++++++++===EXIT===++++++++++=======")
		close(quit)
	}()

	wg.Wait()
	<-quit

	log.Println("microservice exited properly")
}
