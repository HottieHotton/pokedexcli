package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type pokemonDetail struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int    `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func catchCommand(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemon := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	var data pokemonDetail
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	if rand.Intn(data.BaseExperience) >= data.BaseExperience/2 {
		fmt.Println("Oh no! The Pokemon escaped!")
		return nil
	}

	fmt.Printf("%s was caught!\n", data.Name)
	c.pokedex[data.Name] = data
	return nil
}
