package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hottiehotton/pokedexcli/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	commands map[string]cliCommand
	next     *string
	previous *string
	cache    *pokecache.Cache
	pokedex  map[string]pokemonDetail
}



func cleanInput(text string, cfg *config) {
	words := strings.Fields(text)
	if len(words) == 0 {
		return
	}
	cmd := words[0]
	args := words[1:]
	if _, exists := cfg.commands[cmd]; !exists {
		fmt.Printf("Unknown command: %s\n", cmd)
		return
	}
	cfg.commands[cmd].callback(cfg, args)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		cleanInput(input, cfg)
	}
}