package main

import (
	"time"

	"github.com/darkwulf-T/bootdev_pokedexcli/pokeapi"
)

func main() {
	con := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 10*time.Second), Pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(con)

}
