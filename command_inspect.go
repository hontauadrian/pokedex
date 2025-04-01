package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]

	if pokemon, found := cfg.pokedex[name]; found {
		fmt.Printf(
			"Name: %v\n"+
				"Height: %v\n"+
				"Weight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("\t- %s\n", t.Type.Name)
		}
		return nil
	} else {
		return errors.New("you did not yet catch this pokemon\n")
	}
}
