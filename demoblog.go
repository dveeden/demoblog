package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	Comment string
	Author  string
}

type Post struct {
	Id       uint64
	Title    string
	Summary  string
	Content  string
	Comments []Comment
	Rendered string
}

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/posts", postPage)

	go ticker()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
