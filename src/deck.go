package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type called 'deck'
// Which is based on a slice
type deck []string

// Creates a new deck
// Iterating over suits and values.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"clubs", "spades", "Diamonds", "Hearts",
	}

	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jockey",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := fmt.Sprintf("%v of %v", value, suit)
			cards = append(cards, card)
		}
	}
	return cards
}

// Prints every single card on the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Printf("%s\n", card)
	}
}

func deal(d deck, h int) (deck, deck) {
	return d[:h], d[h:]
}

func (d deck) toString() string {
	return strings.Join(d,",\n")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}