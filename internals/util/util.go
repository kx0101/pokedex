package util

import (
	"encoding/json"
	"fmt"
	"github.com/kx0101/pokedex/api"
	"github.com/kx0101/pokedex/internals/shared"
)

func FetchLocations(url string) error {
	if url == "" {
		url = shared.CurrentLocationURL
	}

	entry, exists := shared.PokeCache.Get(url)

	if exists {
		var cachedResults []api.Location
		err := json.Unmarshal(entry, &cachedResults)

		if err == nil {
			PrintLocations(cachedResults)
			return nil
		}
	}

	response, err := api.FetchLocations(url)
	if err != nil {
		return fmt.Errorf("error while fetching locations: %d", err)
	}

	PrintLocations(response.Results)

	responseData, err := json.Marshal(response.Results)
	if err != nil {
		fmt.Println("error while marshaling results of locations.")
	}

	shared.PokeCache.Add(url, responseData)

	shared.NextLocationURL = response.Next
	shared.PrevLocationURL = response.Previous

	return nil
}

func PrintLocations(locations []api.Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
