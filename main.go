package main

import (
	"bitcoin-like-validator/config"
	"bitcoin-like-validator/handler"
	"bitcoin-like-validator/pkg/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Load()

	srv, cleanup := http.ProvideHttpServer(http.ServerConfig{
		Port: 39999,
	})

	han := handler.ProvideHandler(cfg)

	srv.Engine().GET("/submit-answer/:answer", han.SubmitSuccessTask)
	srv.Engine().GET("/:data", han.GetTheHashBase64)
	srv.Engine().GET("/round-has-winner", han.CheckRoundWinner)

	srv.Serve()

	// Gracefully Shutdown
	// Make channel listen for signals from OS
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop

	cleanup()
}
