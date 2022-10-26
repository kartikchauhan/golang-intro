package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.twitter.com",
		"http://www.flipkart.com",
	}

	ch := make(chan string)

	for i, url := range urls {
		fmt.Println("call to " + strconv.Itoa(i) + " starts")
		go checkStatus(url, ch)
		fmt.Println("call to", i, "ends")
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		c <- "Seems like " + url + " is down"
		return
	}

	c <- url + " is up"
}
