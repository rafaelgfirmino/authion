package server

import (
	"context"
	"fmt"
	"github.com/oftall/authion/infra/configuration"
	"github.com/oftall/authion/infra/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	logger := log.New(os.Stdout, "server ", log.LstdFlags)

	listenAddr := fmt.Sprintf(":%v", configuration.Env.Get("server.port"))

	server := http.Server{
		Addr:         listenAddr,
		Handler:      router.Router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("server is ready to handle requests at", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("server stopped")
}