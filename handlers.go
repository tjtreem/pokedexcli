package main

import (
    "fmt"
    "github.com/tjtreem/pokedexcli/internal/pokeapi"
)


func handleMap(config *Config) {
    
    var url string

    if config.Next == nil {
	url = "https://pokeapi.co/api/v2/location-area/"
    } else {	
	url = *config.Next	
    }

    res, err := pokeapi.FetchLocationAreas(url)

    if err != nil {
	fmt.Println("Error:", err)
	return
    }

    for _, area := range res.Results {	
	fmt.Println(area.Name)
    }

    config.Next = res.Next
    config.Previous = res.Previous

}


func handleMapb(config *Config) {

    if config.Previous == nil {
	fmt.Println("you're on the first page")
	return
    }

    url := *config.Previous
	res, err := pokeapi.FetchLocationAreas(url)
	if err != nil {
	    fmt.Println("Error:", err)
	    return
    }

    for _, area := range res.Results {
	fmt.Println(area.Name)
    }

    config.Next = res.Next
    config.Previous = res.Previous
}

