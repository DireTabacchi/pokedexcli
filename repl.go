package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/DireTabacchi/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
    name        string
    description string
    callback    func(*dexState, ...string) error
}

type dexState struct {
    pokeapiClient   pokeapi.Client
    mapState        mapState
    pokedex         map[string]pokeapi.Pokemon
}

func startRepl(ds *dexState) {
    reader := bufio.NewScanner(os.Stdin)
    for {
        prompt()
        reader.Scan()

        comArgs := normalizeInput(reader.Text())
        if len(comArgs) == 0 {
            continue
        }

        comName := comArgs[0]

        command, exists := getCommands()[comName]
        if exists {
            err := command.callback(ds, comArgs[1:]...)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Printf("Unknown command '%s'. Type 'help' for a list of commands.\n", comName)
            continue
        }
    }
}

func normalizeInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

func prompt() {
    fmt.Print("Pokedex > ")
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
        "map": {
            name: "map",
            description: "List a page of locations in the world. Repeated use gets more pages.",
            callback: commandMap,
        },
        "mapb": {
            name: "mapb",
            description: "List the previous page of locations in the world. Repeated use goes further back.",
            callback: commandMapB,
        },
        "explore": {
            name: "explore",
            description: "explore <area_name>; List the Pokemon found in area_name.",
            callback: commandExplore,
        },
        "catch" : {
            name: "catch",
            description: "catch <pokemon_name>; Attempt to cath the Pokemon pokemon_name.",
            callback: commandCatch,
        },
        "inspect": {
            name: "inspect",
            description: "inspect <pokemon_name>; View the stats of the caught Pokemon pokemon_name.",
            callback: commandInspect,
        },
        "pokedex": {
            name: "pokedex",
            description: "Show all Pokemon that have been caught & added to the Pokedex.",
            callback: commandPokedex,
        },
    }
}
