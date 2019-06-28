package main

import (
	"github.com/ataleksand/go-tin/config"
	"github.com/ataleksand/go-tin/storage"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	storage, err := storage.New()
	if err != nil {
		panic(err)
	}
	defer storage.Close()
}
