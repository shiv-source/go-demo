package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://twitter.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinks(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinks(link, c)
		}(l)
	}
}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
	}

	fmt.Println(link, "is up!")
	c <- link
}
