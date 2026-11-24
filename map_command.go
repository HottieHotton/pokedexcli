package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func mapCommand(c *config, args []string) error {
	if c.next == nil {
		fmt.Println("you're on the last page")
		return nil
	}

	url := *c.next
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

func fetchLocation(c *config, url string) ([]byte, error) {
	if cached, ok := c.cache.Get(url); ok {
		return cached, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	c.cache.Add(url, body)
	return body, nil
}
