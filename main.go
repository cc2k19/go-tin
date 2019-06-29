package main

import (
	"github.com/cc2k19/go-tin/config"
	"github.com/cc2k19/go-tin/storage"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	storage, err := storage.New(cfg.Storage)
	if err != nil {
		panic(err)
	}
	defer storage.Close()
}
