package main

import (
	"github.com/cc2k19/go-tin/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	if err := cfg.Validate(); err != nil {
		panic(err)
	}
}
