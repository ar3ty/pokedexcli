package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationList(passedURL *string) (locations, error) {
	url := baseURL + "/location-area"
	if passedURL != nil {
		url = *passedURL
	}

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
