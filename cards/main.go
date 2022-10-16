package main

var decksize int

func main() {
	// cards := deck{"Ace of Spades", "Eight of Heart"}
	// cards = append(cards, "Three of Club")

	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	cards.print()
}
