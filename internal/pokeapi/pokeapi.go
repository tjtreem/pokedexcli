package pokeapi

import (
    "encoding/json"
    "net/http"

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

func FetchLocationAreas(url string) (*LocationAreaResponse, error) {
    res, err := http.Get(url)
    if err != nil {
	return nil, err
    }
    defer res.Body.Close()

    var result LocationAreaResponse
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
	return nil, err
    }

    return &result, nil

}


















