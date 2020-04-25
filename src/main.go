package main

import "fmt"

var card = "Ace of Spades"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println()
	fmt.Println("The actual deck")
	fmt.Printf("%+v", cards)
}
