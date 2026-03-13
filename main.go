package main

import (
	"log"
	"fmt"
	"github.com/alphatee/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	programState := &state {
		cfg: &cfg,
	}	
}
