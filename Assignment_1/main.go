package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range numbers {

		if i%2 == 0 {

			fmt.Println("even")

		} else {

			fmt.Println("odd")

		}

	}

}
