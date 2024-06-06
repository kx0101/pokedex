package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kx0101/pokedex/api"
	cache "github.com/kx0101/pokedex/internals/cache"
)

var (
	currentLocationURL = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
	prevLocationURL    = ""
	nextLocationURL    = ""
	pokecache          = cache.NewCache(time.Second * 5)
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
	err := fetchLocations(nextLocationURL)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	return nil
}

func commandBack() error {
	if prevLocationURL == "" {
		fmt.Println("no previous locations available.")
		return nil
	}

	err := fetchLocations(prevLocationURL)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	return nil
}

func fetchLocations(url string) error {
	if url == "" {
		url = currentLocationURL
	}

	entry, exists := pokecache.Get(url)

	if exists {
		var cachedResults []api.Location
		err := json.Unmarshal(entry, &cachedResults)

		if err == nil {
			printLocations(cachedResults)
			return nil
		}
	}

	response, err := api.FetchLocations(url)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	printLocations(response.Results)

	responseData, err := json.Marshal(response.Results)
	if err != nil {
		fmt.Println("error while marshaling results of locations.")
	}

	pokecache.Add(url, responseData)

	nextLocationURL = response.Next
	prevLocationURL = response.Previous

	return nil
}

func printLocations(locations []api.Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
