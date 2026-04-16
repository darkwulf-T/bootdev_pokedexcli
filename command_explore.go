package main

import "fmt"

func commandExplore(con *config, areaName string) error {
	response, err := con.pokeapiClient.PokemonResponse(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	for _, n := range response.PokemonEncounters {
		fmt.Printf("- %s\n", n.Pokemon.Name)
	}
	return nil
}
