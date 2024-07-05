package main

import (
    "errors"
    "fmt"
    "math/rand"
)

func commandCatch(ds *dexState, args ...string) error {
    if len(args) < 1 {
        return errors.New("No Pokemon name given.")
    } else if len(args) > 1 {
        return errors.New("Too many Pokemon names given.")
    }

    pokemon, err := ds.pokeapiClient.GetPokemon(args[0])
    if err != nil {
        return err
    }

    res := rand.Intn(pokemon.BaseExp)

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
    if res > 40 {
        fmt.Printf("%s escaped!\n", pokemon.Name)
        return nil
    }

    fmt.Printf("%s was caught!\n", pokemon.Name)

    if _, exists := ds.pokedex[pokemon.Name]; !exists {
        ds.pokedex[pokemon.Name] = pokemon
        fmt.Printf("%s has been added to the Pokedex...\n", pokemon.Name)
    }

    return nil
}
