package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocations(url string) (locations, error) {
	locs := locations{}

	res, err := http.Get(url)
	if err != nil {
		return locs, fmt.Errorf("response is not received: %w", err)
	}
	if res.StatusCode > 299 {
		return locs, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locs, fmt.Errorf("reading body is failed: %w", err)
	}

	err = json.Unmarshal(body, &locs)
	if err != nil {
		return locs, fmt.Errorf("json unmarshaling is failed: %w", err)
	}

	return locs, nil
}

func commandMap(cfg *config) error {
	var directionURL string

	if cfg.Backward {
		directionURL = cfg.Previous
	} else {
		directionURL = cfg.Next
	}

	locs, err := getLocations(directionURL)
	if err != nil {
		return err
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	if locs.Previous == nil {
		cfg.Previous = "https://pokeapi.co/api/v2/location-area"
	} else {
		cfg.Previous = locs.Previous.(string)
	}
	cfg.Next = locs.Next

	cfg.Backward = false

	return nil
}

func commandMapB(cfg *config) error {
	cfg.Backward = true
	return commandMap(cfg)
}
