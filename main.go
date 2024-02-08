package main

import "github.com/bphicke/pokedexcli/internal/pokeapi"

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
