package pricefeed

import (
	"context"
	"log"
	"os"
	"os/signal"
	"price-aggregator/pkg/config"
	"syscall"
	"time"
)

// Start initializes the scheduler for fetching prices at regular intervals.
func Start(store *PricesStore, cfg *config.Config) (context.Context, func()) {
	tickerDuration := time.Duration(cfg.FetchSecondsDuration) * time.Second
	ticker := time.NewTicker(tickerDuration)
	ctx, cancel := context.WithCancel(context.Background())

	errChan := make(chan error)

	go func() {
		for err := range errChan {
			log.Println("Error fetching prices:", err)
		}
	}()

	// Set up a channel to listen for interrupt signal (Ctrl+C) or SIGTERM (program termination)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ticker.C:
				go FetchPrices(store, errChan, cfg)
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-sigs:
				log.Println("Received shutdown signal")
				cancel()
			}
		}
	}()

	return ctx, func() {
		cancel()
		close(errChan)
	}
}
