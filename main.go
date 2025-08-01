package main

import (
	"bufio"
	"os"
	"fmt"
	"time"
	"math/rand"
	"github.com/tjtreem/pokedexcli/input"
	"github.com/tjtreem/pokedexcli/internal/pokecache"
)

var commands = map[string]CliCommand{}

func main() {

    rand.Seed(time.Now().UnixNano())
    
    config := &Config{
	pokedex: make(map[string]Pokemon),
    	Cache: pokecache.NewCache(5 * time.Minute),
    }

   
    commands := GetCommands(config)


	scanner := bufio.NewScanner(os.Stdin)
	for {
	    fmt.Print("Pokedex > ")
	    scanner.Scan()
	    rawInput := scanner.Text()

	    words := input.CleanInput(rawInput)

	    if len(words) == 0 {
		continue
	    }
	    
	    cmdName := words[0] 
	    cmd, ok := commands[cmdName]
	    if !ok {
		fmt.Println("Unknown command")
		continue
	    }

	    if cmd.name == "explore" {
		if len(words) < 2 {
		    fmt.Println("You must provide an area name")
		    continue
		}
	    	area_name := words[1]

	    	err := commandExplore(config, area_name)
	    	if err != nil {
		    fmt.Println(err)
	    	}
	    } else if cmd.name == "catch" {
	    	if len(words) < 2 {
		    fmt.Println("please provide a pokemon name")
		    continue
		}
		pokemonName := words[1]
		
		err := commandCatch(config, pokemonName)
		if err != nil {
		    fmt.Println(err)
		}
	    } else if cmd.name == "inspect" {
		if len(words) < 2 {
		    fmt.Println("please provide a pokemon name")
		    continue
		}
		pokemonName := words[1]

		err := commandInspect(config, pokemonName)
		if err != nil {
		    fmt.Println(err)
		}
	    } else {
		err := cmd.callback()
		if err != nil {
		    fmt.Println(err)
		}   
	    }
		
	}
}

