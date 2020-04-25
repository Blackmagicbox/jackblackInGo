package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got: %d instead \n", len(d))
	}

	if d[0].getName() != "Ace of Spades" {
		t.Errorf("Expected the first card of the deck to be an Ace of Spades but got: %s instead\n", d[0].getName())
	}

	if d[len(d) - 1].getName() != "King of Clubs" {
		t.Errorf("Expected the last card of the deck to be a King of Clubs but got: %s", d[len(d) - 1].getName())
	}
}

// TODO: Refactor newDeckFromFile method to support new struct
//func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
//	os.Remove("_decktesting")
//
//	deck := newDeck()
//
//	deck.saveToFile("_decktesting")
//
//	loadedDeck := newDeckFromFile("_decktesting")
//
//	if len(loadedDeck) != 52 {
//		t.Errorf("Expected deck length of 52, but got: %d instead \n", len(loadedDeck))
//	}
//
//	os.Remove("_decktesting")
//}

