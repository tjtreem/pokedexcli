package main

import (
	"github.com/tjtreem/pokedexcli/internal/pokecache"
)


type Config struct {
	pokedex		map[string]Pokemon
	Next		*string
	Previous	*string
	Cache		*pokecache.Cache

}


type CliCommand struct {
	name		string
	description	string
	callback	func() error

}


type LocationAreaResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		} `json:"pokemon_encounters"`
}


type StatDetail struct {
	Name		string	`json:"name"`
}

type Stat struct {
	BaseStat	int		`json:"base_stat"`
	Stat		StatDetail	`json:"stat"`
}


type TypeDetail struct {
	Name		string		`json:"name"`
}

type Type struct {
	Type		TypeDetail	`json:"type"`
}


type Pokemon struct {
	Name 		string `json:"name"`
	BaseExperience	int	`json:"base_experience"`
	Height		int	`json:"height"`
	Weight		int	`json:"weight"`
	Stats		[]Stat	`json:"stats"`
	Types		[]Type	`json:"types"`

}


