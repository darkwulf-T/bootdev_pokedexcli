package main

import "fmt"

func commandMapf(con *config, secondarg string) error {
	response, err := con.pokeapiClient.RequestFunction(con.next)
	if err != nil {
		return err
	}
	con.next = response.Next
	con.previous = response.Previous
	for _, n := range response.Results {
		fmt.Println(n.Name)
	}
	return nil
}

func commandMapb(con *config, secondarg string) error {
	if con.previous == nil {
		return fmt.Errorf("This is the first page")
	}
	response, err := con.pokeapiClient.RequestFunction(con.previous)
	if err != nil {
		return err
	}
	con.next = response.Next
	con.previous = response.Previous
	for _, n := range response.Results {
		fmt.Println(n.Name)
	}
	return nil
}
