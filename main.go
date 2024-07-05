package main

import (
	"time"

	"github.com/DireTabacchi/pokedexcli/internal/pokeapi"
)

func main() {
    pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
    ds := dexState{
        pokeapiClient: pokeClient,
        pokedex: make(map[string]pokeapi.Pokemon),
    }

    startRepl(&ds)
}
