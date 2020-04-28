package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.json")
	fmt.Printf("%+v", cards)
}
