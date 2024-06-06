package shared

import (
	"github.com/kx0101/pokedex/internals/cache"
	"time"
)

var (
	CurrentLocationURL = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
	PrevLocationURL    = ""
	NextLocationURL    = ""
	PokeCache          = cache.NewCache(time.Second * 5)
)
