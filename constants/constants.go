package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
)

func constants() {
	fmt.Printf("%vKB\n", KB)
	fmt.Printf("%vMB\n", MB)
}
