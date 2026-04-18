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
		secondarg := ""
		switch firstW {
		case "explore":
			if len(cleanUserI) < 2 {
				fmt.Println("An area must be selected")
				continue
			} else {
				secondarg = cleanUserI[1]
			}
		case "catch":
			if len(cleanUserI) < 2 {
				fmt.Println("An Pokemon must be selected")
				continue
			} else {
				secondarg = cleanUserI[1]
			}
		default:
		}

		cmd, ok := getCommands()[firstW]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(con, secondarg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
