package main

import "fmt"

var decksize int

func main() {
	cards := deck{"Ace of Spades", "Eight of Heart"}
	cards = append(cards, "Three of Club")

	for index, card := range cards {
		fmt.Println(index, card)
	}
}
