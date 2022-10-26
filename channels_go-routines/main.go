package main

import (
	"fmt"
	"net/http"
	"time"
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

	for _, url := range urls {
		go checkStatus(url, ch)
	}

	// for {
	// 	go checkStatus(<-ch, ch)
	// }

	// Alternative for loop

	// Whenever channel ch receives a value, assign it to url and make a call to function checkStatus.
	for url := range ch { // iterate over channel ch.
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, ch)
		}(url)
	}
}

func checkStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Seems like " + url + " is down")
		c <- url
		return
	}

	fmt.Println(url + " is up")
	c <- url
}
