package main

import (
	"bufio"
	"fmt"
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
		areaName := ""
		if firstW == "explore" {
			if len(cleanUserI) < 2 {
				fmt.Println("An area must be selected")
				continue
			} else {
				areaName = cleanUserI[1]
			}
		}

		cmd, ok := getCommands()[firstW]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(con, areaName)
		if err != nil {
			fmt.Println(err)
		}
	}
}
