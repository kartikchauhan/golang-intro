package main

import "fmt"

var decksize int

func main() {
	// cards := deck{"Ace of Spades", "Eight of Heart"}
	// cards = append(cards, "Three of Club")

	cards := newDeck()

	// hand, remainingDeck := cards.deal(5)

	// hand.print()
	// remainingDeck.print()
	// cards.print()

	fmt.Println(cards.toString())

}
