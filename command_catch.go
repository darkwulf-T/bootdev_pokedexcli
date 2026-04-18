package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(con *config, secondarg string) error {
	pokemon, err := con.pokeapiClient.CatchResponse(secondarg)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", secondarg)

	catchChance := rand.IntN(pokemon.BaseExperience)

	if catchChance > 30 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command")
	con.Pokedex[pokemon.Name] = pokemon

	return nil
}
