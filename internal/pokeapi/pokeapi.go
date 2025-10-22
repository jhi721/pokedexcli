package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jhi721/pokedexcli/internal/pokecache"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeApiClient struct {
	cache pokecache.Cache
}

type api interface {
	Get(url string) (Locations, error)
}

func NewPokeApiClient() PokeApiClient {
	return PokeApiClient{
		cache: pokecache.NewCache(5 * time.Second),
	}
}

func (p PokeApiClient) Get(url string) (Locations, error) {
	val, exists := p.cache.Get(url)

	if exists {
		var locations Locations

		if err := json.Unmarshal(val, &locations); err != nil {
			log.Fatal(err)
		}

		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}

	var locations Locations
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, nil
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		log.Fatal(err)
	}

	p.cache.Add(url, data)

	return locations, nil
}
