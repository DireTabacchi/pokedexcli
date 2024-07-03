package main

import "fmt"

func commandHelp(ds *dexState) error {
    fmt.Println("\nWelcome to the Pokedex!")
    fmt.Printf("Usage:\n\n")
    for _, command := range getCommands() {
        fmt.Printf("%v: %v\n", command.name, command.description)
    }
    fmt.Printf("\n")

    return nil
}
