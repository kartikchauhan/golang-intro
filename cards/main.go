package main

var decksize int

func main() {
	cards := newDeckFromFile("deck_1")
	cards.print()
}
