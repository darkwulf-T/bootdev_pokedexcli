package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/darkwulf-T/bootdev_pokedexcli/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(con *config, areaName string) error
}

type config struct {
	next          *string
	previous      *string
	pokeapiClient pokeapi.Client
	Pokedex       map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays location from the pokeapi. Consecutive calls advance up the pages",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays location from the previous page from the pokeapi",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows all Pokemon in a selected area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a selected Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Look up details about a caught Pokemon",
			callback:    commandInspect,
		},
	}
}

func commandExit(con *config, secondarg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(con *config, secondarg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
