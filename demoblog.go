package main

import (
	"embed"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed css
var cssFS embed.FS

var dburi string
var defdburi string = "root@tcp(127.0.0.1:4000)/blog?parseTime=true"

func main() {
	flag.StringVar(&dburi, "db", defdburi, "database URI")
	flag.Parse()

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
