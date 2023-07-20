package main

import (
	"embed"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed css
var cssFS embed.FS

var dburi string = "root@tcp(127.0.0.1:4000)/blog?parseTime=true"

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/posts", postPage)
	http.HandleFunc("/analytics", analyticsPage)
	http.HandleFunc("/likes", likesApi)

	fs := http.FileServer(http.FS(cssFS))
	http.Handle("/css/", fs)

	go ticker()
	go tickerChecker()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
