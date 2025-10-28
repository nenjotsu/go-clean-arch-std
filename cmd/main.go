package main

import (
	"log"
	"net/http"
	"time"

	"github.com/nenjotsu/go-clean-arch-std/internal/app/handlers"
	"github.com/nenjotsu/go-clean-arch-std/internal/app/middlewares"
)

func main() {
	mux := http.NewServeMux()
	handlers.RegisterHandlers(mux)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      middlewares.LoggingMiddleware(mux), // from middleware.go
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Running on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
