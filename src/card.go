package main

type Card struct {
	suit string
	printName string
	name string
	value int
}

func (c Card) getName() string {
	return c.printName
}