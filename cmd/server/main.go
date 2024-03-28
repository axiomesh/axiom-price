package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"price-aggregator/internal/api"
	"price-aggregator/internal/pricefeed"
	"price-aggregator/pkg/config"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	store := pricefeed.NewPricesStore()
	// Initialize price fetch scheduler
	ctx, stop := pricefeed.Start(store, cfg)

	httpPort := ":" + cfg.HTTPPort
	srv := http.Server{Addr: httpPort}

	// Setup HTTP server and route, passing the store to the handler
	http.HandleFunc("/api/ticker/price", func(w http.ResponseWriter, r *http.Request) {
		api.HandleTickerPrice(w, r, store)
	})

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// Start the HTTP server
	log.Printf("HTTP server started on %v", httpPort)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server Shutdown: %v", err)
	}

	stop()

	log.Println("Server exited")
}
