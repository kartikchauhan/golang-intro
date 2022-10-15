package main

var decksize int

func main() {
	cards := deck{"Ace of Spades", "Eight of Heart"}
	cards = append(cards, "Three of Club")

	cards.print()
}
