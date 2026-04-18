package main

import "fmt"

func commandInspect(con *config, secondarg string) error {
	pok, ok := con.Pokedex[secondarg]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pok.Name)
	fmt.Printf("Height: %v\n", pok.Height)
	fmt.Printf("Weight: %v\n", pok.Weight)
	fmt.Println("Stats:")
	for _, stat := range pok.Stats {
		fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pok.Types {
		fmt.Printf("   -%s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(con *config, secondarg string) error {
	fmt.Println("Your Pokedex:")
	for p := range con.Pokedex {
		fmt.Printf("   -%s\n", p)
	}
	return nil
}
