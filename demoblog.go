package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed css
var cssFS embed.FS

var dburi string
var defdburi string = "root@tcp(127.0.0.1:4000)/blog?parseTime=true"
var bpm int

func main() {
	flag.StringVar(&dburi, "db", defdburi, "database URI")
	flag.IntVar(&bpm, "bpm", 60, "beats-per-minute for ticker")
	flag.Parse()

	if !strings.Contains(dburi, "parseTime=true") {
		panic("Please add 'parseTime=true' to the dburi")
	}

	if strings.Contains(dburi, "tidbcloud") {
		if !strings.Contains(dburi, "tls=true") {
			log.Print("Please add 'tls=true' to the dburi when connecting to TiDB Cloud")
		}
	}

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/posts", postPage)
	http.HandleFunc("/analytics", analyticsPage)
	http.HandleFunc("/likes", likesApi)

	fs := http.FileServer(http.FS(cssFS))
	http.Handle("/css/", fs)

	// 60/bpm results in 0, which means no sleep.
	if bpm > 60 {
		log.Print("Limiting ticker to max 60 BPM")
		bpm = 60
	}

	if bpm < 1 {
		log.Print("Ticker disabled")
	} else {
		go ticker()
		go tickerChecker()
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
