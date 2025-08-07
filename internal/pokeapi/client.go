package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ar3ty/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
	Cache      pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) getResponse(url string) ([]byte, error) {
	cached, isPresent := c.Cache.Get(url)

	if isPresent {
		return cached, nil
	}

	body := []byte{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, fmt.Errorf("request is not formed: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return body, fmt.Errorf("response is not received: %w", err)
	}
	if res.StatusCode > 299 {
		return body, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return body, fmt.Errorf("reading body is failed: %w", err)
	}

	c.Cache.Add(url, body)

	return body, nil
}

func (c *Client) GetLocationList(passedURL *string) (locations, error) {
	url := baseURL + "/location-area"
	if passedURL != nil {
		url = *passedURL
	}

	locs := locations{}

	body, err := c.getResponse(url)
	if err != nil {
		return locs, err
	}

	err = json.Unmarshal(body, &locs)
	if err != nil {
		return locs, fmt.Errorf("json unmarshaling is failed: %w", err)
	}

	return locs, nil
}

func (c *Client) GetPokemonList(area string) (pokencounters, error) {
	url := baseURL + "/location-area/" + area

	pocs := pokencounters{}

	body, err := c.getResponse(url)
	if err != nil {
		return pocs, err
	}

	err = json.Unmarshal(body, &pocs)
	if err != nil {
		return pocs, fmt.Errorf("json unmarshaling is failed: %w", err)
	}

	return pocs, nil
}

func (c *Client) GetPokemon(target string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + target

	pok := Pokemon{}

	body, err := c.getResponse(url)
	if err != nil {
		return pok, err
	}

	err = json.Unmarshal(body, &pok)
	if err != nil {
		return pok, fmt.Errorf("json unmarshaling is failed: %w", err)
	}

	return pok, nil
}
