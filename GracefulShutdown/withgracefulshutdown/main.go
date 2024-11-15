package main

import (
	"context"
	"log"
	"net/http"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
	"github.com/quii/go-graceful-shutdown/acceptancetests"
)

func main() {
	var (
		ctx        = context.Background()
		httpServer = &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}
		server     = gracefulshutdown.NewServer(httpServer)
	)
	if err := server.ListenAndServe(ctx); err != nil {
		log.Fatalf("failed to shut down gracefully, some responses may have been lost: %v", err)
	}
	log.Println("graceful shutdown completed. All responses accounted for:")
}
