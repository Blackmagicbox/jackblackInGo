package main

type card struct {
	suit string `json:"suit"`
	printName string `json:"print_name"`
	name string `json:"name"`
	value int `json:"value"`
}

func (c card) getName() string {
	return c.printName
}