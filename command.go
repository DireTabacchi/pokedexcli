package main

import (
    "fmt"
    "os"
)

type cliCommand struct {
    name string
    description string
    callback func() error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Show the help message",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
    }
}

func commandHelp() error {
    fmt.Println("\nWelcome to the Pokedex!")
    fmt.Printf("Usage:\n\n")
    for _, command := range getCommands() {
        fmt.Printf("%v: %v\n", command.name, command.description)
    }
    fmt.Printf("\n")

    return nil
}

func commandExit() error {
    os.Exit(0)
    return nil
}
