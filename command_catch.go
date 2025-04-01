package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	if _, found := cfg.pokedex[name]; found {
		fmt.Printf("%s is already inside de pokedex\n", name)
		return nil
	}

	for {
		fmt.Printf("Throwing a Pokeball at %s...\n", name)
		if rand.Intn(pokemonResp.BaseExperience) > pokemonResp.BaseExperience/2 {
			cfg.pokedex[name] = pokemonResp
			fmt.Printf("%s was caught!\n", name)
			return nil
		} else {
			fmt.Printf("%s escaped!\n", name)
		}
	}
}
