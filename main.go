package main

import (
	"time"

	"github.com/mtgrinstead/pokedexcli/internal/pokeapi"
)

type config struct {
    pokeapiClient       pokeapi.Client
    nextLocationAreaURL *string
    prevLocationAreaURL *string
    caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
    cfg := config{
      pokeapiClient: pokeapi.NewClient(time.Hour),
      caughtPokemon: make(map[string]pokeapi.Pokemon), 
    }
    startRepl(&cfg)
}
