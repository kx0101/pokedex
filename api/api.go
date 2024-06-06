package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kx0101/pokedex/internals/shared"
)

type Location struct {
	Name string
}

type LocationsResponse struct {
	Next      string     `json:"next"`
	Previous  string     `json:"previous"`
	Locations []Location `json:"results"`
}

type Stats struct {
	Base   int `json:"base_stat"`
	Effort int `json:"effort"`
	Stat   struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Types struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type PokemonStats struct {
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func fetch(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading from body: %v", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	return nil
}

func FetchLocations(url string) (LocationsResponse, error) {
	var result LocationsResponse

	if url == "" {
		url = shared.CurrentLocationURL
	}

	entry, exists := shared.PokeCache.Get(url)
	if exists {
		var cachedResults LocationsResponse
		err := json.Unmarshal(entry, &cachedResults)

		if err == nil {
			shared.NextLocationURL = cachedResults.Next
			shared.PrevLocationURL = cachedResults.Previous

			return LocationsResponse{
				Locations: cachedResults.Locations,
			}, nil
		}
	}

	err := fetch(url, &result)
	responseData, errParse := json.Marshal(result)
	if errParse != nil {
		return result, fmt.Errorf("error while marshaling results of locations: %s", err)
	}

	shared.PokeCache.Add(url, responseData)

	shared.NextLocationURL = result.Next
	shared.PrevLocationURL = result.Previous

	return result, err
}

func FetchLocation(location string) (LocationResponse, error) {
	var result LocationResponse

	entry, exists := shared.PokeCache.Get(location)
	if exists {
		var cachedPokemonEncounters []PokemonEncounter
		err := json.Unmarshal(entry, &cachedPokemonEncounters)

		if err == nil {
			return LocationResponse{
				PokemonEncounters: cachedPokemonEncounters,
			}, nil
		}
	}

	err := fetch(shared.GetPokemonsFromLocationURL+location, &result)
	responseData, errParse := json.Marshal(result)
	if errParse != nil {
		return result, fmt.Errorf("error while marshaling pokemon encounters: %s", err)
	}

	shared.PokeCache.Add(location, responseData)

	return result, err
}

func FetchPokemon(pokemon string) (PokemonStats, error) {
	var result PokemonStats

	entry, exists := shared.PokeCache.Get(pokemon)
	if exists {
		var cachedPokemonStats PokemonStats
		err := json.Unmarshal(entry, &cachedPokemonStats)

		if err == nil {
			return cachedPokemonStats, nil
		}
	}

	err := fetch(shared.GetPokemonDataURL+pokemon, &result)
	if err != nil {
		return result, fmt.Errorf("error while fetching for pokemon data: %s", err)
	}

	responseData, errParse := json.Marshal(result)
	if errParse != nil {
		return result, fmt.Errorf("error while marshaling pokemon data: %s", errParse)
	}

	shared.PokeCache.Add(pokemon, responseData)

	return result, nil
}
