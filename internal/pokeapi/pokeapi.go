package pokeapi

import (
    "encoding/json"
    "net/http"
    "io"
    "github.com/tjtreem/pokedexcli/internal/pokecache"
)


type LocationArea struct {
    Name string `json:"name"`
    URL  string `json:"url"`

}

type LocationAreaResponse struct {
    Count	int		`json:"count"`
    Next	*string		`json:"next"`
    Previous	*string		`json:"previous"`
    Results	[]LocationArea	`json:"results"`

}


// Fetches location areas from the provided url

func FetchLocationAreas(url string, cache *pokecache.Cache) (*LocationAreaResponse, error) {
    data, ok := cache.Get(url)
    if ok {
	var result LocationAreaResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
	    return nil, err
	}
	return &result, nil
    }
    res, err := http.Get(url)
    if err != nil {
	return nil, err
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
	return nil, err
    }

    var result LocationAreaResponse
    err = json.Unmarshal(body, &result)
    if err != nil {
	return nil, err
    }

    cache.Add(url, body)

    return &result, nil

}


















