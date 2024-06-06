package commands

import (
	"fmt"
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
			Description: "Explore a specific area for pokemons",
			Callback:    commandExplore,
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
