package main

import (
	"fmt"
	"net/http"
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

	for {
		go checkLinks(<-c, c)  // Infinite time it will ping the list of sites
		
	}
}

func checkLinks(link string, c chan string ) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
	return
}
