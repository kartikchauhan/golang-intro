package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func printEndpoint(s string) {
	fmt.Printf("Endpoint => %v\n", s)
}

func getTabs(res http.ResponseWriter, req *http.Request) {
	printEndpoint("/tabs")

	data := []byte("List of all tabs")
	res.Write(data)
}

func getTab(res http.ResponseWriter, req *http.Request) {
	printEndpoint("/tabs/{tab_id}")

	data := []byte("Tab id 2")
	res.Write(data)
}

func main() {
	// Setup handler for endpoints

	http.HandleFunc("/tabs", getTabs)
	http.HandleFunc("/tabs/:tab_id", getTab)

	// Start HTTP server
	// TODO => Add handler in ListenAndServe which will print the endpoint
	err := http.ListenAndServe(":3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server closed ", err)
	} else if err != nil {
		log.Fatal("Error while starting the server ", err)
	}
}
