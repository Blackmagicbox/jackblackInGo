package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type called 'deck'
// Which is based on a slice

type deck []card

// Creates a new deck Iterating over suits and values.
func newDeck() deck {
	cards := deck{}

	cardNames := []string {
		"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jockey", "Queen", "King",
	}

	cardSuits := []string{
		"Spades", "Hearts", "Diamonds", "Clubs",
	}

	cardValues := []int{ 14,2,3,4,5,6,7,8,9,10,11,12,13}

	for _, suit := range cardSuits {
		for _, name := range cardNames {
			cardName := fmt.Sprintf("%v of %v", name, suit)
			card := card{ suit: suit, printName: cardName, name: name }
			cards = append(cards, card)
		}
	}

	for k, value := range cardValues {
		cards[k].value = value
	}
	return cards
}

// Iterate over all deck Items printing each one

func (d deck) print() {
	for _, card := range d {
		fmt.Printf("%s\n", card.getName())
	}
}

// Deal cards based on a hand size
// It basically splits the deck in 2; hand and remaining

func deal(d deck, h int) (deck, deck) {
	return d[:h], d[h:]
}

// Converts the deck from Slice to string coma separated

func (d deck) toString() string {
	var cardNames []string

	for _, card := range d {
		cardNames = append(cardNames, card.getName())
	}

	return strings.Join(cardNames, ",")
}

// Save the deck converted to string into a file.
func (d deck) saveToFile(fileName string) error {
	file, _ := json.Marshal([]card(d))

	return ioutil.WriteFile(fileName, file, 0644)
}

// Create a new deck from a coma-separated string
// TODO: Refactor this method to create a deck from Json files.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	d := deck{}

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	_ = json.Unmarshal(bs, &d)

	return d
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
