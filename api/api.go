package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string
}

type LocationResponse struct {
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func FetchLocations(url string) (LocationResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error fetching data: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return LocationResponse{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error reading from body: %v", err)
	}

	var result LocationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return result, nil
}
