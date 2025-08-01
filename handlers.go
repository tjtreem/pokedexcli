package main

import (
    "encoding/json"
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

    res, err := pokeapi.FetchLocationAreas(url, config.Cache)

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
	res, err := pokeapi.FetchLocationAreas(url, config.Cache)
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


func handlePokemon(data []byte) error {
	var response LocationAreaResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
	    return fmt.Errorf("failed to unmarshal location area data: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range response.PokemonEncounters {
	    fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}


func catchPokemon(config *Config, args []string) error {
	if len(args) == 0 {
	    return fmt.Errorf("please provide a pokemon name")
	}
	return commandCatch(config, args[0])
}


func inspectPokemon(config *Config, args []string) error {
	if len(args) == 0 {
	    return fmt.Errorf("please provide a pokemon name")
	}
	return commandInspect(config, args[0])
}


func handlePokedex(config *Config) error {
	if len(config.pokedex) == 0 {
	    fmt.Println("Your Pokedex is empty")
	    return nil
	}

	fmt.Println("Your Pokedex:")

	for name := range config.pokedex {
	    fmt.Println(" -", name)
	}

	return nil
}















