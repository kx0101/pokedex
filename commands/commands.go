package commands

import (
	"fmt"
	"os"
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
