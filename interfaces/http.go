package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom Writer
type writerLog struct{}

func request() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	wl := writerLog{}

	io.Copy(wl, resp.Body)
}

// Custom writer function
func (writerLog) Write(bs []byte) (int, error) {
	fmt.Printf("[%v byte] %v", len(bs), string(bs))

	return len(bs), nil
}
