package repl

import "os"

func commandExit(store *Store, args ...string) error {
	os.Exit(0)
	return nil
}
