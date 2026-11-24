package main

import (
	"time"

	"github.com/hottiehotton/pokedexcli/pokecache"
)

func main() {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	cfg := &config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    exitCommand,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    helpCommand,
			},
			"map": {
				name:        "map",
				description: "Go to the next location area batch results",
				callback:    mapCommand,
			},
			"mapb": {
				name:        "mapb",
				description: "Go to the previous location area batch results",
				callback:    mapbCommand,
			},
			"explore": {
				name:        "explore",
				description: "Explore the current location requested",
				callback:    exploreCommand,
			},
			"catch": {
				name:        "catch",
				description: "Attempt to catch a Pokemon",
				callback:    catchCommand,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a caught Pokemon",
				callback:    inspectCommand,
			},
			"pokedex": {
				name:        "pokedex",
				description: "View your caught Pokemon",
				callback:    pokedexCommand,
			},
		},
		next:     &baseURL,
		previous: nil,
		cache:    pokecache.NewCache(5 * time.Minute),
		pokedex:  make(map[string]pokemonDetail),
	}
	startRepl(cfg)
}