package main

import "fmt"

type Card struct {
	suit string
	name string
	value int
}

func (c Card) getName() string {
	return c.name
}

func (c Card) print() {
	fmt.Printf("%s\n,", c.name)
}