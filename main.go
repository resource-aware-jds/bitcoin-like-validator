package main

import (
	"bitcoin-like-validator/pkg/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv, cleanup := http.ProvideHttpServer(http.ServerConfig{
		Port: 39999,
	})

	srv.Serve()

	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop

	cleanup()
}
