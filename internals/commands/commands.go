package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/kx0101/pokedex/api"
	"github.com/kx0101/pokedex/internals/shared"
	"github.com/kx0101/pokedex/internals/util"
)

type ClipCommand struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

var Commands = map[string]ClipCommand{}
var Pokedex = map[string]api.PokemonStats{}

func InitCommands() {
	Commands = map[string]ClipCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Fetch next 20 Pokemon locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "map",
			Description: "Fetch previous 20 Pokemon locations",
			Callback:    commandBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a specific area for pokemons: explore <area-name>",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "cache",
			Description: "Attempt to catch a specific pokemon: catch <pokemon-name>",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect your pokemon: inspect <pokemon-name>",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "View your pokedex",
			Callback:    commandPokedex,
		},
	}
}

func commandHelp(args ...string) error {
	fmt.Println("\nWelcome to Pokedex!")
	fmt.Println("Available commands: ")

	fmt.Println()
	for name, cmd := range Commands {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}

	fmt.Println()
	return nil
}

func commandExit(args ...string) error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(1)
	return nil
}

func commandMap(args ...string) error {
	err := util.FindLocations(shared.NextLocationURL)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	return nil
}

func commandBack(args ...string) error {
	if shared.PrevLocationURL == "" {
		fmt.Println("no previous locations available.")
		return nil
	}

	err := util.FindLocations(shared.PrevLocationURL)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	return nil
}

func commandExplore(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to provide an area name: explore <area-name>")
	}

	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	location := api.Location{
		Name: areaName,
	}

	util.ExploreLocation(location)
	return nil
}

func commandCatch(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to provide a pokemon name: catch <pokemon-name>")
	}

	pokemonName := args[0]
	_, exists := Pokedex[pokemonName]
	if exists {
		fmt.Printf("You have already caught: %s\n", pokemonName)
		return nil
	}

	pokemonStats, err := util.FindPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error while fetching for pokemon data: %s", err)
	}

	fmt.Printf("Throwing a pokeball at %s...\n", pokemonName)

	randomNumber := rand.Float64()
	if randomNumber > (1.0 / float64(pokemonStats.BaseExperience)) {
		Pokedex[pokemonName] = pokemonStats
		fmt.Printf("%s was caught\n", pokemonName)
	} else {
		fmt.Printf("pokemon %s escaped\n", pokemonName)
		return nil
	}

	return nil
}

func commandInspect(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to provide a pokemon name: inspect <pokemon-name>")
	}

	pokemonName := args[0]
	pokemonStats, exists := Pokedex[pokemonName]
	if !exists {
		fmt.Printf("You haven't caught %s yet\n", pokemonName)
		return nil
	}

	prettyFormat, err := json.MarshalIndent(pokemonStats, "", " ")
	if err != nil {
		return fmt.Errorf("error while marshaling pokemon data: %s", err)
	}

	fmt.Println(string(prettyFormat))
	return nil
}

func commandPokedex(args ...string) error {
	for pokemon := range Pokedex {
		fmt.Printf("\n- %s", pokemon)
	}

	fmt.Println()
	return nil
}
