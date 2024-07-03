package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
)

type locationList struct {
    Next string `json:"next"`
    Previous string `json:"previous"`
    Locations []location `json:"results"`
}

type location struct {
    Name string
    Url string
}

func (ll locationList) String() string {
    var llString string
    llString += fmt.Sprintln("locationList{")

    llString += fmt.Sprintln("Next:", ll.Next)
    llString += fmt.Sprintln("Previous:", ll.Previous)
    llString += fmt.Sprintln("Locations:", ll.Locations)

    llString += fmt.Sprintln("}")
    
    return llString
}

type mapState struct {
    Next string
    Previous string
}

func commandMap(ds *dexState) error {
    var res *http.Response
    var err error
    if ds.mapState.Next == "" {
        res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
    } else {
        res, err = http.Get(ds.mapState.Next)
    }
    if err != nil {
        fmt.Println("Error occurred with GET")
        return err
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
        fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        return errors.New(fmt.Sprintf("Request returned status code: %d\n", res.StatusCode))
    }
    if err != nil {
        fmt.Println("Error reading body")
        return err
    }

    var locations locationList
    err = json.Unmarshal(body, &locations)
    if err != nil {
        fmt.Println("Error parsing location body")
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
    if ds.mapState.Previous == "" {
        fmt.Println("Nothing to go back to!")
        return nil
    }

    var res *http.Response
    var err error
    res, err = http.Get(ds.mapState.Previous)

    if err != nil {
        fmt.Println("Error occurred with GET")
        return err
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
        fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        return errors.New(fmt.Sprintf("Request returned status code: %d\n", res.StatusCode))
    }
    if err != nil {
        fmt.Println("Error reading body")
        return err
    }

    var locations locationList
    err = json.Unmarshal(body, &locations)
    if err != nil {
        fmt.Println("Error parsing location body")
        return err
    }

    ds.mapState.Next = locations.Next
    ds.mapState.Previous = locations.Previous

    for _, loc := range locations.Locations {
        fmt.Println(loc.Name)
    }

    return nil
}

func printResponse(res *http.Response) {
        fmt.Println("Response:")
        fmt.Println("  Status:", res.Status)
        fmt.Println("  StatusCode:", res.StatusCode)
        fmt.Println("  Proto:", res.Proto)
        fmt.Println("    ProtoMajor:", res.ProtoMajor)
        fmt.Println("    ProtoMinor:", res.ProtoMinor)
        fmt.Println("  Header:", res.Header)
        fmt.Println("  Body:", res.Body)
        fmt.Println("  ContentLength:", res.ContentLength)
        fmt.Println("  TransferEncoding:", res.TransferEncoding)
        fmt.Println("  Close:", res.Close)
        fmt.Println("  Uncompressed:", res.Uncompressed)
        fmt.Println("  Trailer:", res.Trailer)
        fmt.Println("  Request:", res.Request)
        fmt.Println("  TLS:", res.TLS)

}
