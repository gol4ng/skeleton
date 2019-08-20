package main

import (
	"log"
	"net/http"
	"os"
	"syscall"

	"google.golang.org/grpc"

	"github.com/gol4ng/signal"

	grpcSrv "github.com/gol4ng/skeleton/cmd/server/grpc"
	httpSrv "github.com/gol4ng/skeleton/cmd/server/http"
	"github.com/gol4ng/skeleton/config"
	"github.com/gol4ng/skeleton/internal/service"
)

func main() {
	// loading config
	cfg := config.NewServer()

	// creating your service container
	container := service.NewContainer(cfg)

	// init servers
	httpServer := httpSrv.NewServer(container)
	grpcServer := grpcSrv.NewServer(container)

	idleConnsClosed := make(chan struct{})

	// gracefully stop app on interrupt signal
	// kill app when receiving kill signal or 2x interrupt signal (ex: 2x ctr+c)
	defer signal.SubscribeWithKiller(func(signal os.Signal) {
		log.Println(signal.String(), "signal received : gracefully stopping application")
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
	}, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP)()

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
