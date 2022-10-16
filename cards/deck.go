package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardValue := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, cardValue+" of "+suit)
		}
	}

	return cards
}

// Golang's standard suggests to use the first letter of type as the variable name
func (d deck) print() { // print is a receiver which receives type deck

	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
