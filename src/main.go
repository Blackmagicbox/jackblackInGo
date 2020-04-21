package main

var card = "Ace of Spades"

func main() {
	cards := newDeckFromFile("my_cards.txt")
	cards.print()
}
