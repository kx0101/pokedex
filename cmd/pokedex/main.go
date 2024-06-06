package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kx0101/pokedex/internals/commands"
)

func main() {
	commands.InitCommands()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println()
	fmt.Println("Enter a command (type 'help' for a list of commands)")

	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}

		parts := strings.Fields(input)
		cmdName := parts[0]
		args := parts[1:]

		cmd, exists := commands.Commands[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %s", cmdName)
			fmt.Println()
			fmt.Println("Enter a command (type 'help' for a list of commands)")
			continue
		}

		if err := cmd.Callback(args...); err != nil {
			fmt.Printf("Error executing command '%s': %v\n", cmdName, err)
		}

		fmt.Println("Enter a command (type 'help' for a list of commands):")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
