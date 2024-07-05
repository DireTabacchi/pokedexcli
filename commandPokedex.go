package main

import (
    "errors"
    "fmt"
)

func commandPokedex(ds *dexState, args ...string) error {
    if len(ds.pokedex) < 1 {
        return errors.New("You haven't caught any Pokemon yet.")
    }

    pokemonString := "Your Pokedex:\n"

    for pokemonName := range ds.pokedex {
        pokemonString += fmt.Sprintf("  - %s\n", pokemonName)
    }

    fmt.Print(pokemonString)

    return nil
}
