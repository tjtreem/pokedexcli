package main

import (
	"bufio"
	"os"
	"fmt"
	"github.com/tjtreem/pokedexcli/input"
)

type Config struct {
	Next		*string
	Previous	*string

}


type cliCommand struct {
	name		string
	description	string
	callback	func() error

}


func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp() error {
    fmt.Println(`Welcome to the Pokedex!
Usage:
`)
    for _, cmd := range commands {
	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

var commands = map[string]cliCommand{}

func main() {
    
    config := &Config{}

    commands = map[string]cliCommand {
    "help": {
    name:		"help",
    description:	"Displays a help message",
    callback:		commandHelp,
    },
    "exit": {
    name:		"exit",
    description:	"Exit the Pokedex",
    callback:		commandExit,
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

	scanner := bufio.NewScanner(os.Stdin)
	for {
	    fmt.Print("Pokedex > ")
	    scanner.Scan()
	    rawInput := scanner.Text()

	    words := input.CleanInput(rawInput)

	    if len(words) == 0 {
		continue
	    }
	    
	 
	    cmd, ok := commands[words[0]]
	    if ok {
		err := cmd.callback()
		if err != nil {
		    fmt.Println(err)
		}
	    } else {
		fmt.Println("Unknown command")
	    }


	    
	    
	}
		
}

