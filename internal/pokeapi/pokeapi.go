package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jhi721/pokedexcli/internal/pokecache"
)

type Locations struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonInformation struct {
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Name                   string `json:"name"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Weight int `json:"weight"`
}

type PokeApiClient struct {
	cache pokecache.Cache
}

func NewPokeApiClient() PokeApiClient {
	return PokeApiClient{
		cache: pokecache.NewCache(5 * time.Second),
	}
}

func (p PokeApiClient) GetPokemonByName(name string) (PokemonInformation, error) {
	val, exists := p.cache.Get(name)

	if exists {
		var pokemonInformation PokemonInformation

		if err := json.Unmarshal(val, &pokemonInformation); err != nil {
			return PokemonInformation{}, fmt.Errorf("error unmarshaling cached entry %v", err)
		}

		return pokemonInformation, nil
	}

	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		return PokemonInformation{}, fmt.Errorf("error getting %s pokemon entry", name)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return PokemonInformation{}, fmt.Errorf("cannot find pokemon with the name %s", name)
	} else if res.StatusCode > 299 {
		return PokemonInformation{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	var pokemonInformation PokemonInformation
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonInformation, fmt.Errorf("error reading from the body %w", err)
	}

	if err := json.Unmarshal(data, &pokemonInformation); err != nil {
		return PokemonInformation{}, fmt.Errorf("error unmarshaling entry from the api %w", err)
	}

	p.cache.Add(name, data)

	return pokemonInformation, nil
}

func (p PokeApiClient) GetByName(name string) (LocationArea, error) {
	val, exists := p.cache.Get(name)

	if exists {
		var locationArea LocationArea

		if err := json.Unmarshal(val, &locationArea); err != nil {
			return LocationArea{}, fmt.Errorf("error unmarshaling cached entry %v", err)
		}

		return locationArea, nil
	}

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + name)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error getting %s location area entry", name)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return LocationArea{}, fmt.Errorf("cannot find area with the name %s", name)
	} else if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	var locationArea LocationArea
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading from the body %w", err)
	}

	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("error unmarshaling entry from the api %v", err)
	}

	p.cache.Add(name, data)

	return locationArea, nil
}

func (p PokeApiClient) Get(url string) (Locations, error) {
	val, exists := p.cache.Get(url)

	if exists {
		var locations Locations

		if err := json.Unmarshal(val, &locations); err != nil {
			return Locations{}, fmt.Errorf("error unmarshaling cached entry %v", err)
		}

		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Locations{}, fmt.Errorf("error getting location area entries")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Locations{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	var locations Locations
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, fmt.Errorf("error reading from the body %w", err)
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return Locations{}, fmt.Errorf("error unmarshaling entry from the api %v", err)
	}

	p.cache.Add(url, data)

	return locations, nil
}
