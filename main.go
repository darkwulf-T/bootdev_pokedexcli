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
		fmt.Printf("Your command was: %s\n", firstW)
	}
}