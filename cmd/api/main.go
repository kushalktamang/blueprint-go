package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kushalktamang/blueprint-go/internal/config"
	"github.com/kushalktamang/blueprint-go/internal/handlers"
)

func main() {

	cfg := config.MustLoad()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.HealthCheck)

	server := http.Server{
		Addr:         ":" + cfg.PORT,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("server is listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
