package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PostResponse struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const API = "https://jsonplaceholder.typicode.com/posts"

func getAllPost(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(API)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var posts []PostResponse
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(body)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/posts", getAllPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
