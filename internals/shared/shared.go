package shared

import (
	"github.com/kx0101/pokedex/internals/cache"
	"time"
)

var (
	CurrentLocationURL         = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	PrevLocationURL            = ""
	NextLocationURL            = ""
	GetPokemonsFromLocationURL = "https://pokeapi.co/api/v2/location-area/"
	GetPokemonDataURL          = "https://pokeapi.co/api/v2/pokemon/"
	PokeCache                  = cache.NewCache(time.Second * 5)
)
