package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.json")
	cards.print()
	//fmt.Println()
	//fmt.Println("The actual deck")
	//fmt.Printf("%+v", cards)
}
