package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userI := scanner.Text()
		cleanUserI := cleanInput(userI)
		firstW := cleanUserI[0]
		
		val, ok := getCommands()[firstW]
		if !ok {
			fmt.Println("Unknown command")
		}
		val.callback()
	}
}