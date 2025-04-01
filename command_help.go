package main

import "fmt"

func commandHelp(c *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, supportedCommand := range getCommands() {
		fmt.Printf("%s: %s\n", name, supportedCommand.description)
	}
	fmt.Println()
	return nil
}
