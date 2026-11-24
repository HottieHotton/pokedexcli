package main

import (
	"encoding/json"
	"fmt"
)

func mapbCommand(c *config, args []string) error {
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *c.previous
	body, err := fetchLocation(c, url)
	if err != nil {
		return err
	}

	var data locationAreaResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	c.next = data.Next
	c.previous = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
