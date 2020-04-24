package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type called 'deck'
// Which is based on a slice

type deck []string

// Creates a new deck
// Iterating over suits and values.

func newDeck() deck {
	cards := deck{}

	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jockey", "Queen", "King",
	}

	cardSuits := []string{
		"Spades", "Hearts", "Diamonds", "Clubs",
	}


	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := fmt.Sprintf("%v of %v", value, suit)
			cards = append(cards, card)
		}
	}
	return cards
}

// Iterate over all deck Items printing each one

func (d deck) print() {
	for _, card := range d {
		fmt.Printf("%s\n", card)
	}
}

// Deal cards based on a hand size
// It basically splits the deck in 2; hand and remaining

func deal(d deck, h int) (deck, deck) {
	return d[:h], d[h:]
}

// Converts the deck from Slice to string coma separated

func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Save the deck converted to string into a file.

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Create a new deck from a coma-separated string

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

// Shuffles the deck each time the function is called
// Generating a new Seed each time.

func (d deck) shuffleCards() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
