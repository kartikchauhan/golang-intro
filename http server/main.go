package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"context"
)

const keyServerAddr = "serverAddress"

func printEndpoint(s string) {
	fmt.Printf("Endpoint => %v\n", s)
}

func getTabs(res http.ResponseWriter, req *http.Request) {
	printEndpoint("/tabs")

	ctx := req.Context()

	fmt.Printf("%s \n", ctx.Value(keyServerAddr))

	data := []byte("List of all tabs")
	res.Write(data)
}

func getTab(res http.ResponseWriter, req *http.Request) {
	printEndpoint("/tabs/{tab_id}")

	ctx := req.Context()
	
	fmt.Printf("%s \n", ctx.Value(keyServerAddr))

	data := []byte("Tab id 2")
	res.Write(data)
}

func main() {
	// Setup handler for endpoints

	mux := http.NewServeMux()
	mux.HandleFunc("/tabs", getTabs)
	mux.HandleFunc("/tabs/:tab_id", getTab)

	// Start HTTP server
	// TODO => Add handler in ListenAndServe which will print the endpoint
	err := http.ListenAndServe(":3000", mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server closed ", err)
	} else if err != nil {
		log.Fatal("Error while starting the server ", err)
	}
}
