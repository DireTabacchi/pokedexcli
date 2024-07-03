package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func startRepl() {
    reader := bufio.NewScanner(os.Stdin)
    for {
        prompt()
        reader.Scan()
        // READ
        comArgs := normalizeInput(reader.Text())

        if len(comArgs) == 0 {
            continue
        }

        comName := comArgs[0]

        command, exists := getCommands()[comName]

        if exists {
            err := command.callback()
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Printf("'%s' is not a command. Type 'help' for a list of commands.\n", comName)
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
