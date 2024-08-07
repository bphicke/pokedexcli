package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, t := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", t.Stat.Name, t.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s \n", t.Type.Name)
	}
	return nil
}
