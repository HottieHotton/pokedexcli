package main

import (
	"fmt"
)

func inspectCommand(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemon := args[0]
	detail, exists := c.pokedex[pokemon]
	if !exists {
		return fmt.Errorf("pokemon %s not found in your Pokedex", pokemon)
	}
	fmt.Printf("Name: %s\n", detail.Name)
	fmt.Printf("Base Experience: %d\n", detail.BaseExperience)
	fmt.Printf("Height: %d\n", detail.Height)
	fmt.Printf("Weight: %d\n", detail.Weight)
	fmt.Println("Stats:")
	for _, stat := range detail.Stats {
		fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, t := range detail.Types {
		fmt.Printf("Type: %s\n", t.Type.Name)
	}
	return nil
}