package util

import (
	"fmt"

	"github.com/kx0101/pokedex/api"
)

func FindLocations(url string) error {
	response, err := api.FetchLocations(url)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	PrintLocations(response.Locations)
	return nil
}

func ExploreLocation(location api.Location) error {
	pokemonEncounters, err := api.FetchLocation(location.Name)
	if err != nil {
		return fmt.Errorf("error while fetching for pokemons of the location: %s", err)
	}

	PrintPokemonNames(pokemonEncounters.PokemonEncounters)
	return nil
}

func FindPokemon(pokemonName string) (api.PokemonStats, error) {
	pokemonStats, err := api.FetchPokemon(pokemonName)
	if err != nil {
		return api.PokemonStats{}, fmt.Errorf("error while fetching for data of pokemon: %s", err)
	}

	return pokemonStats, nil
}

func PrintLocations(locations []api.Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func PrintPokemonNames(pokemonEncounters []api.PokemonEncounter) {
	fmt.Println("Found Pokemon:")
	fmt.Println()

	for _, pokemonEncounter := range pokemonEncounters {
		fmt.Printf("\n- %s", pokemonEncounter.Pokemon.Name)
	}

	fmt.Println()
	fmt.Println()
}
