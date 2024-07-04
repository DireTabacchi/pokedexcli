package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

)

type locationList struct {
    Next *string `json:"next"`
    Previous *string `json:"previous"`
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

func (c *Client) ListLocations(pageURL *string) (locationList, error) {
    endpoint := baseURL + "/location-area"
    if pageURL != nil {
        endpoint = *pageURL
    }

    if val, ok := c.cache.Get(endpoint); ok {
        locations := locationList{}
        err := json.Unmarshal(val, &locations)
        if err != nil {
            return locationList{}, err
        }

        return locations, nil
    }

    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        return locationList{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return locationList{}, err
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return locationList{}, err
    }

    locations := locationList{}
    err = json.Unmarshal(data, &locations)
    if err != nil {
        return locationList{}, err
    }

    c.cache.Add(endpoint, data)

    return locations, nil

}
