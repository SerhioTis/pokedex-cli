package repl

import (
	"errors"
	"fmt"
)

func commandMapf(store *Store, args ...string) error {
	locationsResp, err := store.PokeaApiClient.GetLocationList(store.nextLocationsURL)
	if err != nil {
		return err
	}

	store.nextLocationsURL = locationsResp.Next
	store.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(store *Store, args ...string) error {
	if store.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := store.PokeaApiClient.GetLocationList(store.prevLocationsURL)
	if err != nil {
		return err
	}

	store.nextLocationsURL = locationResp.Next
	store.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
