package main

import "fmt"

type deck []string

// Golang's standard suggests to use the first letter of type as the variable name 
func (d deck) print() { // print is a receiver which receives type deck

	for index, card := range d {
		fmt.Println(index, card)
	}
}
