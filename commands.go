package main

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
	"os"
	"math/rand"
)


func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp(commands map[string]CliCommand) error {
    fmt.Println(`Welcome to the Pokedex!
Usage:`)
    for _, cmd := range commands {
	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

func commandExplore(config *Config, area_name string) error {
    var url string
    url = "https://pokeapi.co/api/v2/location-area/" + area_name
   
    data, ok := config.Cache.Get(url)
    if ok {
	fmt.Printf("Exploring %s (from cache)...\n", area_name)
	
	err := handlePokemon(data)
	if err != nil {
	    return err
	}
	return nil
    } else {
	resp, err := http.Get(url)
	if err != nil {
	    return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
    	    if err != nil {
	    return err
    	}	    
    	config.Cache.Add(url, body)
	
	err = handlePokemon(body)
	if err != nil {
	    return err
	}
	return nil
    }
    return nil
}


func commandCatch(config *Config, pokemonName string) error {
    url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

    resp, err := http.Get(url)
    if err != nil {
	return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
	return fmt.Errorf("pokemon not found")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
	return err
    }

    var pokemon Pokemon
    err = json.Unmarshal(body, &pokemon)
    if err != nil {
	return err
    }

    catchChance := 100 - (pokemon.BaseExperience / 3)
    if catchChance < 10 {
	catchChance = 10
    }

    randomRoll := rand.Intn(100)
    if randomRoll < catchChance {
	fmt.Printf("%s was caught!\n", pokemonName)
	config.pokedex[pokemonName] = pokemon
    } else {
	fmt.Printf("%s escaped!\n", pokemonName)
    }
    return nil
}

func commandInspect(config *Config, pokemonName string) error {
    pokemon, exists := config.pokedex[pokemonName]
    if !exists {
	fmt.Println("you have not caught that pokemon")
	return nil
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)

    fmt.Printf("Stats:\n")
    for _, stat := range pokemon.Stats {
	    fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Printf("Types:\n")
    for _, pokemonType := range pokemon.Types {
	fmt.Printf(" - %s\n", pokemonType.Type.Name)
    }

    return nil
}


func commandPokedex(config *Config) error {
    return handlePokedex(config)
}


func GetCommands(config *Config) map[string]CliCommand {
    return map[string]CliCommand{
    	"help": {
    	name:		"help",
    	description:	"Displays a help message",
    	callback: func() error {
	    return commandHelp(GetCommands(config))
    	},
	},	
    	"exit": {
    	name:		"exit",
    	description:	"Exit the Pokedex",
    	callback:		commandExit,
    	},
    	"explore": {
    	name:		"explore",
    	description:	"Explore the location area",
    	callback: func() error {
            return nil
    	},
    	},
    	"catch": {
    	name:		"catch",
    	description:	"Attempt to catch a pokemon",
    	callback: func() error {
	    return nil
    	},
    	},
    	"inspect": {
    	name:		"inspect",
    	description:	"inspect captured pokemon",
    	callback: func() error {
	    return nil
    	},
    	},
    	"pokedex": {
    	name:		"pokedex",
    	description:	"Lists your caught Pokemon",
    	callback: func() error {
	    return commandPokedex(config)
    	},
    	},
    	"map": {
	name:		"map",
	description:	"Show the next 20 location areas",
	callback: func() error {
	    handleMap(config)
	    return nil
	},
    	},
    	"mapb": {
	name:		"mapb",
	description:	"Show previous 20 location areas",
	callback: func() error {
	    handleMapb(config)
	    return nil
	},
    	},
    }
}
