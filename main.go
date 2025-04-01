package main

import (
	"github.com/hontauadrian/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
