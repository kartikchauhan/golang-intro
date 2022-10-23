package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { // Assumption: printGreeting has common code
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string { // Assumption: getGreeting has very custom code
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
