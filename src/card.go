package main

import "fmt"

type Card struct {
	suit string
	printName string
	name string
	value int
}

func (c Card) getName() string {
	return c.printName
}

func (c Card) print() {
	fmt.Printf("%s\n,", c.printName)
}