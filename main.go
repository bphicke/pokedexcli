package main

import (
	"time"

	"github.com/bphicke/pokedexcli/internal/pokeapi"
)

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
