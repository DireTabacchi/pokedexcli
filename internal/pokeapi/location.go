package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"

)

func (c *Client) GetLocation(locName *string) (locationDetailed, error) {
    endpoint := baseURL + "/location-area/" + *locName

    if val, ok := c.cache.Get(endpoint); ok {
        location := locationDetailed{}
        err := json.Unmarshal(val, &location)
        if err != nil {
            return locationDetailed{}, err
        }

        return location, nil
    }

    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        return locationDetailed{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return locationDetailed{}, err
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return locationDetailed{}, err
    }

    location := locationDetailed{}
    err = json.Unmarshal(data, &location)
    if err != nil {
        return locationDetailed{}, err
    }

    c.cache.Add(endpoint, data)

    return location, nil
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
