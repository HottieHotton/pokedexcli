package main

import (
	"encoding/json"
	"fmt"
)

type locationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func exploreCommand(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area name")
	}
	location := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + location

	fmt.Printf("Exploring %s...\n", location)

	body, err := fetchLocation(c, url)
	if err != nil {
		return err
	}

	var data locationAreaDetail
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, enc := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
