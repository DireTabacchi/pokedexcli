package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokename string) (Pokemon, error) {
    endpoint := baseURL + "/pokemon/" + pokename

    if val, ok := c.cache.Get(endpoint); ok {
        mon := Pokemon{}
        err := json.Unmarshal(val, &mon)
        if err != nil {
            return Pokemon{}, err
        }
        return mon, nil
    }

    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        return Pokemon{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, err
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return Pokemon{}, nil
    }

    mon := Pokemon{}
    err = json.Unmarshal(data, &mon)
    if err != nil {
        return Pokemon{}, err
    }

    c.cache.Add(endpoint, data)

    return mon, nil
}
