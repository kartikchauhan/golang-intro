package main

import (
	"fmt"
	"net/http"
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
		// fmt.Println("call to " + strconv.Itoa(i) + " starts")
		go checkStatus(url, ch)
		// fmt.Println("call to", i, "ends")
	}

	for {
		go checkStatus(<-ch, ch)
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
