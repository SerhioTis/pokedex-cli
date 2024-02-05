package repl

import (
	"fmt"
)

func commandPokedex(store *Store, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range store.CaughtPokemon {
		fmt.Println(" -", p.Name)
	}

	return nil
}
