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
		cache: pokecache.NewCache(5 * time.Minute),
	}
}

func checkCache[T Locations | LocationArea | PokemonInformation](cache pokecache.Cache, key string) (T, bool, error) {
	var cachedVal T

	val, exists := cache.Get(key)

	if !exists {
		return cachedVal, exists, nil
	}

	if err := json.Unmarshal(val, &cachedVal); err != nil {
		return cachedVal, exists, fmt.Errorf("error unmarshaling cached entry %w", err)
	}

	return cachedVal, exists, nil
}

func getRequest[T Locations | LocationArea | PokemonInformation](cache pokecache.Cache, url string) (T, error) {
	var payload T

	res, err := http.Get(url)
	if err != nil {
		return payload, fmt.Errorf("error GET %s\n%w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return payload, fmt.Errorf("GET %s not found", url)
	} else if res.StatusCode > 299 {
		return payload, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return payload, fmt.Errorf("error reading from the body %w", err)
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		return payload, fmt.Errorf("error unmarshaling entry from the api %w", err)
	}

	cache.Add(url, data)

	return payload, nil
}

func (p PokeApiClient) GetPokemonByName(name string) (PokemonInformation, error) {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/" + name

	pokemonInformationFromCache, exists, err := checkCache[PokemonInformation](p.cache, baseUrl)
	if err != nil {
		return PokemonInformation{}, err
	}
	if exists {
		return pokemonInformationFromCache, nil
	}

	pokemonInformationFromApi, err := getRequest[PokemonInformation](p.cache, baseUrl)
	if err != nil {
		return PokemonInformation{}, err
	}

	return pokemonInformationFromApi, nil
}

func (p PokeApiClient) GetLocationAreaByName(name string) (LocationArea, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/" + name

	locationAreaFromCache, exists, err := checkCache[LocationArea](p.cache, baseUrl)
	if err != nil {
		return LocationArea{}, err
	}
	if exists {
		return locationAreaFromCache, nil
	}

	locationAreaFromApi, err := getRequest[LocationArea](p.cache, baseUrl)
	if err != nil {
		return LocationArea{}, err
	}

	return locationAreaFromApi, nil

}

func (p PokeApiClient) GetLocations(url string) (Locations, error) {
	locationsFromCache, exists, err := checkCache[Locations](p.cache, url)
	if err != nil {
		return Locations{}, err
	}
	if exists {
		return locationsFromCache, nil
	}

	locationsFromApi, err := getRequest[Locations](p.cache, url)
	if err != nil {
		return Locations{}, err
	}

	return locationsFromApi, nil
}
