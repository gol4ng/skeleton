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
	// example loading config
	cfg := config.NewServer()

	// creating your service container
	container := service.NewContainer(cfg)

	// init servers
	httpServer := httpSrv.NewServer(container)
	grpcServer := grpcSrv.NewServer(container)

	// signal handler
	idleConnsClosed := make(chan struct{})
	stopping := false
	defer signal.Subscribe(func(signal os.Signal) {
		// kill application if os.kill signal is received or if user interrupt graceful shutdown (i.e hit ctr+c two times)
		if stopping || signal == os.Kill {
			log.Println("killing application")
			os.Exit(0)
		}
		// graceful shutdow
		stopping = true
		log.Println(signal.String(), "signal received : gracefully stopping application")
		println("Press `ctrl+c` again to kill application.")
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
