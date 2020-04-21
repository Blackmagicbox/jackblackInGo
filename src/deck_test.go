package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got: %d instead \n", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card of the deck to be an Ace of Spades but got: %s instead\n", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected the last card of the deck to be a King of Clubs but got: %s", d[len(d) - 1])
	}
}
