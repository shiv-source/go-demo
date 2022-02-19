package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")

	if err != nil {
		fmt.Println(err)
	}

	out, err := os.Create("data.txt")

	if err != nil {
		fmt.Println(err)

	}
	io.Copy(out, resp.Body)
}
