package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		saveToJson(comments, "comments.json")

	}
}

func saveToJson(data []Comment, fileName string) {

	commentsJson, _ := json.Marshal(data)
	var out bytes.Buffer
	json.Indent(&out, commentsJson, "", "\t")

	err := ioutil.WriteFile(fileName, out.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data successfully added")
	}
}

func main() {

	client = &http.Client{Timeout: 10 * time.Second}
	getComment()
}
