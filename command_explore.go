package main

import "fmt"

func commandExplore(con *config, secondarg string) error {
	response, err := con.pokeapiClient.PokemonResponse(secondarg)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", secondarg)
	for _, n := range response.PokemonEncounters {
		fmt.Printf("- %s\n", n.Pokemon.Name)
	}
	return nil
}
