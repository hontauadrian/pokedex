package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	locationAreaResp, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", name)
	for _, loc := range locationAreaResp.PokemonEncounters {
		fmt.Printf("-%s\n", loc.Pokemon.Name)
	}
	return nil
}
