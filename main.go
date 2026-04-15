package main

import (
	"github.com/darkwulf-T/bootdev_pokedexcli/pokeapi"
	"time"
)

func main() {
	con := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second, 10 * time.Second),
	}
	startRepl(con)

}