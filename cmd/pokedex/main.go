package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kx0101/pokedex/internals/commands"
)

func main() {
	commands.InitCommands()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a command (type 'help' for a list of commands)")

	for scanner.Scan() {
		cmdName := scanner.Text()
		if len(cmdName) == 0 {
			continue
		}

		cmd, exists := commands.Commands[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %s", cmdName)
			fmt.Println()
			fmt.Println("Enter a command (type 'help' for a list of commands)")
			continue
		}

		if err := cmd.Callback(); err != nil {
			fmt.Printf("Error executing command '%s': %v\n", cmdName, err)
		}

		fmt.Println("Enter a command (type 'help' for a list of commands):")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
