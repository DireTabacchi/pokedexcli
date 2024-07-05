package pokeapi

import "fmt"

type locationList struct {
    Next *string `json:"next"`
    Previous *string `json:"previous"`
    Locations []location `json:"results"`
}

type location struct {
    Name string
    Url string
}

type locationDetailed struct {
    Name string
    PokemonEncounters []struct{
        Pokemon apiResource `json:"pokemon"`
    } `json:"pokemon_encounters"`
}

func (ll locationList) String() string {
    var llString string
    llString += fmt.Sprintln("locationList{", ll.Next, ll.Previous, ll.Locations, "}")
    
    return llString
}

