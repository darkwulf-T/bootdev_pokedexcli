package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(con *config, secondarg string) error {
	pokemon, err := con.pokeapiClient.CatchResponse(secondarg)
	if err != nil {
		if err == fmt.Errorf("Error with the GET request: Not Found") {
			return fmt.Errorf("The choosen Pokemon either don't exist or is not available. Please select an existing Pokemon.")
		} else {
			return err
		}
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", secondarg)

	catchChance := rand.IntN(pokemon.BaseExperience)

	if catchChance > 30 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
