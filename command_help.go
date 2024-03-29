package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the pokedex help menu!")
	fmt.Println("Here are your available commands:")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s ", cmd.name, cmd.description)
		fmt.Println("")
	}
	return nil
}
