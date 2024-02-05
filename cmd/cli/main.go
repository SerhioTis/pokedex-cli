package main

import (
	"time"

	"github.com/SerhioTis/pokedex-cli/internal/pokeapi"
	"github.com/SerhioTis/pokedex-cli/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	store := &repl.Store{
		CaughtPokemon:  map[string]pokeapi.Pokemon{},
		PokeaApiClient: pokeClient,
	}

	repl.StartRepl(store)
}
