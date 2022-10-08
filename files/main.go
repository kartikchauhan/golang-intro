package main

import (
	"fmt"
	"log"
	"os"
)

func fileRead() {
	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal("Error while opening file", err)

	data := make([]byte, 100)

	count, err := file.Read(data)

	fmt.Printf("[%v bytes] %s", count, data)
}

func fileWrite() {
	file, err := os.Create("sample2.txt")

	if err != nil {
		log.Fatal("Error while creating file", err)
	}

	data := []byte("Sample content")

	count, err := file.Write(data)

	if err != nil {
		log.Fatal("Error while writing to file", err)
	}

	fmt.Printf("[%v bytes] %s", count, data)
}

func main() {
	// fileRead()

	// fileWrite()
}
