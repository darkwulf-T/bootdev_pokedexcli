package main

import (
	"fmt"
	"bufio"
	"os"
)

func startRepl(con *config) {
		scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userI := scanner.Text()
		cleanUserI := cleanInput(userI)
		firstW := cleanUserI[0]
		
		cmd, ok := getCommands()[firstW]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(con)
		if err != nil {
			fmt.Println(err)
		}
	}
}
