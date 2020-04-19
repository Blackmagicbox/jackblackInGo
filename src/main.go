package main

import "fmt"

var card = "Ace of Spades"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Your hand:")
	hand.print()
	fmt.Println()
	fmt.Println("Remaining Cards")
	remainingCards.print()
	fmt.Println(hand.toString())
}
