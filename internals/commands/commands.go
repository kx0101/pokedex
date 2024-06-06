package commands

import (
	"fmt"
	"os"

	"github.com/kx0101/pokedex/internals/shared"
	"github.com/kx0101/pokedex/internals/util"
)

type ClipCommand struct {
	Name        string
	Description string
	Callback    func() error
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
			Description: "Explore a specific area",
			Callback:    commandExplore,
		},
	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to Pokedex!")
	fmt.Println("Available commands: ")

	fmt.Println()
	for name, cmd := range Commands {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}

	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(1)
	return nil
}

func commandMap() error {
	err := util.FetchLocations(shared.NextLocationURL)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	return nil
}

func commandBack() error {
	if shared.PrevLocationURL == "" {
		fmt.Println("no previous locations available.")
		return nil
	}

	err := util.FetchLocations(shared.PrevLocationURL)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	return nil
}

func commandExplore() error {
	return nil
}
