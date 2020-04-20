package main

import "fmt"

var card = "Ace of Spades"

func main() {
	cards := newDeckFromFile("my_cards.txt")
	shuffledCards := cards.shuffleCards()

	hand, remainingCards := deal(shuffledCards, 5)

	fmt.Println()
	fmt.Println("Remaining Cards")
	remainingCards.print()
	fmt.Println()
	fmt.Println("Shuffled cards")
	shuffledCards.print()
	fmt.Println()
	fmt.Println("Your hand:")
	hand.print()

}
