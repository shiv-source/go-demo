package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var client *http.Client

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func getJson(url string, target interface{}) error {

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func getComment() {
	url := "https://jsonplaceholder.typicode.com/comments"

	var comments []Comment

	err := getJson(url, &comments)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Post Id  =>", comments[0].PostId)
		saveToReadme(comments, "comments.md")

	}
}

func saveToReadme(comments []Comment, fileName string) {

	const (
		header = `# Fetch data to markdown file
Get data from JSON PLACEHOLDER 

| PostId  | Id  | Name | Email | Body |
| ------- | --- | ---- | ----- | ---- |
`
	)

	readme, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0666)

	if err != nil {
		fmt.Println(err)
		os.Create(fileName)
	}
	defer readme.Close()
	readme.WriteString(header)

	for _, comment := range comments {
		
		name := strings.ReplaceAll(comment.Name, "\n" , "")
		body := strings.ReplaceAll(comment.Body, "\n", "")

		readme.WriteString(fmt.Sprintf("| %d | %d | %s | %s | %s |\n", comment.PostId, comment.Id, name, comment.Email, body ))
	}

	fmt.Println("Successfully written the Markdown file")
}

func main() {

	client = &http.Client{Timeout: 10 * time.Second}
	getComment()
}
