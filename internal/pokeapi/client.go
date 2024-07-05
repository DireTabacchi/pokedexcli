package pokeapi

import (
    "net/http"
    "time"

    "github.com/DireTabacchi/pokedexcli/internal/pokecache"
)

type Client struct {
    cache       pokecache.Cache
    httpClient  http.Client
    Pokedex     map[string]Pokemon
}

func NewClient(timeout, cacheInterval time.Duration) Client {
    return Client{
        cache: pokecache.NewCache(cacheInterval),
        httpClient: http.Client{
            Timeout: timeout,
        },
        Pokedex: make(map[string]Pokemon),
    }
}
