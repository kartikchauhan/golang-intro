package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#fff000",
	}

	capitals := make(map[string]string)

	capitals["india"] = "Delhi"
	capitals["australia"] = "canberra"
	capitals["pakistan"] = "islamabad"

	delete(capitals, "pakistan")

	printColors(colors)
}

func printColors(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("%v has hex code %v\n", color, hex)
	}
}
