package main

import (
    "errors"
    "fmt"
)


type mapState struct {
    Next        *string
    Previous    *string
}

func commandMap(ds *dexState) error {
    locations, err := ds.pokeapiClient.ListLocations(ds.mapState.Next)
    if err != nil {
        return err
    }

    ds.mapState.Next = locations.Next
    ds.mapState.Previous = locations.Previous

    for _, loc := range locations.Locations {
        fmt.Println(loc.Name)
    }

    return nil
}

func commandMapB(ds *dexState) error {
    if ds.mapState.Previous == nil {
        return errors.New("Already on the first page!")
    }

    locations, err := ds.pokeapiClient.ListLocations(ds.mapState.Previous)
    if err != nil {
        return err
    }

    ds.mapState.Next = locations.Next
    ds.mapState.Previous = locations.Previous

    for _, loc := range locations.Locations {
        fmt.Println(loc.Name)
    }

    return nil
}
