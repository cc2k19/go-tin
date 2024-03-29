package main

import (
	"context"
	"github.com/cc2k19/go-tin/api"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/cc2k19/go-tin/config"
	"github.com/cc2k19/go-tin/server"
	"github.com/cc2k19/go-tin/storage"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handleInterrupts(ctx, cancel)

	wg := &sync.WaitGroup{}

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	if err = cfg.Validate(); err != nil {
		panic(err)
	}

	db, err := storage.New(cfg.Storage)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := storage.NewRepository(db)

	api := api.New(repository)

	server := server.New(cfg.Server, api)
	server.Run(ctx, wg)

	wg.Wait()
}

// handleInterrupts handles process signal interrupts
func handleInterrupts(ctx context.Context, cancel context.CancelFunc) {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt)
	go func() {
		select {
		case <-term:
			log.Println("Received OS interrupt, exiting gracefully...")
			cancel()
		case <-ctx.Done():
			return
		}
	}()
}
