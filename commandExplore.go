package main

import (
    "errors"
    "fmt"
)

func commandExplore(ds *dexState, args ...string) error {
    if len(args) < 1 {
        return errors.New("No area name given.")
    } else if len(args) > 1 {
        return errors.New("Too many area names given.")
    }

    location, err := ds.pokeapiClient.GetLocation(&args[0])
    if err != nil {
        return err
    }

    fmt.Println("Found pokemon:")

    for _, encounter := range location.PokemonEncounters {
        fmt.Println(" -", encounter.Pokemon.Name)
    }

    return nil
}
