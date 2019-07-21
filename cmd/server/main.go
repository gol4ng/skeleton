package main

import (
	"net/http"
	"os"
	"syscall"

	"google.golang.org/grpc"

	grpcSrv "github.com/gol4ng/skeleton/cmd/server/grpc"
	httpSrv "github.com/gol4ng/skeleton/cmd/server/http"
	"github.com/gol4ng/skeleton/config"
	"github.com/gol4ng/skeleton/internal/service"
	"github.com/gol4ng/skeleton/pkg/signal"
)

func main() {
	// example loading config
	cfg := config.NewServer()

	// creating your service container
	container := service.NewContainer(cfg)

	// init servers
	httpServer := httpSrv.NewServer(container)
	grpcServer := grpcSrv.NewServer(container)

	// signal handler
	idleConnsClosed := make(chan struct{})
	defer signal.Subscribe(func(signal os.Signal) {
		println(signal.String(), "signal received. stopping...")
		if err := httpServer.Stop(); err != nil {
			println(err)
		}
		if err := grpcServer.Stop(); err != nil {
			println(err)
		}
		if err := container.Unload(); err != nil {
			println(err)
		}
		close(idleConnsClosed)
	}, os.Interrupt, os.Kill, syscall.SIGTERM)()

	// starting servers
	go func() {
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	if err := grpcServer.Start(); err != nil && err != grpc.ErrServerStopped {
		panic(err)
	}

	<-idleConnsClosed
}
