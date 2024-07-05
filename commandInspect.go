package main

import (
    "errors"
    "fmt"
)

func commandInspect(ds *dexState, args ...string) error {
    if len(args) < 1 {
        return errors.New("No Pokemon name given.")
    } else if len(args) > 1 {
        return errors.New("Too many Pokemon names given.")
    }
    mon, exists := ds.pokedex[args[0]]
    if !exists {
        return errors.New("You have not caught that Pokemon yet.")
    }

    var inspection string
    inspection += fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\n", mon.Name, mon.Height, mon.Weight)

    inspection += fmt.Sprintln("Stats:")
    for _, stat := range mon.Stats {
        inspection += fmt.Sprintf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
    }

    inspection += fmt.Sprintln("Types:")
    for _, t := range mon.Types {
        inspection += fmt.Sprintf("  - %s\n", t.Type.Name)
    }

    fmt.Print(inspection)

    return nil
}
