package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed css
var cssFS embed.FS

type Comment struct {
	Comment string
	Author  string
	Created time.Time
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
	http.HandleFunc("/analytics", analyticsPage)

	fs := http.FileServer(http.FS(cssFS))
	http.Handle("/css/", fs)

	go ticker()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
