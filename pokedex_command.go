package main

import (
	"fmt"
)

func pokedexCommand(c *config, args []string) error {
	if len(c.pokedex) == 0 {
		return fmt.Errorf("you haven't caught any Pokemon yet")
	}
	fmt.Println("Your caught Pokemon:")
	for _, pokemon := range c.pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}