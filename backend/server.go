package main

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/config"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/database/datastore"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/registry"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const defaultPort = "8080"

func main() {
	err := config.Setup("config/config_dev.yaml")
	if err != nil {
		panic(fmt.Errorf("error getting config %v", err))
	}

	err = datastore.SetupDB()
	if err != nil {
		panic(fmt.Errorf("error connecting to db %v", err))
	}
	listener, errChan := runHTTPServer()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func(errChan chan error) {
		errCh := errChan
		select {
		case <-sigCh:
			_ = listener.Close()
		case _ = <-errCh:
			_ = listener.Close()
		}
		cancel()
	}(errChan)
	<-ctx.Done()
}

func runHTTPServer() (listener net.Listener, ch chan error) {
	router, err := registry.InitHTTPServer(context.Background())
	if err != nil {
		panic(err)
	}

	conf := config.GetConfig()

	addr := fmt.Sprintf(":%s", conf.Server.ServerPort)
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	ch = make(chan error)
	go func() {
		ch <- http.Serve(listener, router)
	}()

	return listener, ch
}
