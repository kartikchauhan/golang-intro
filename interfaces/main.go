package main

import "fmt"

// To whom it may concern, our program has a new type called 'bot'
type bot interface {
	getGreeting() string
	// If you are a type in this program with a function called 'getGreeting' and you return a string then you are now an honorary member of type 'bot'
	// Now that you're also an honorary member of type 'bot', you can now call this function called 'printGreeting'
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	request()
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
