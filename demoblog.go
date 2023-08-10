package main

import (
	"embed"
	"errors"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

var dburi string = "root@tcp(127.0.0.1:4000)/blog?parseTime=true"

//go:embed ui/dist
var uiFiles embed.FS

func main() {
	var initFlag bool

	flag.Func("db", "database URI", func(flagValue string) (err error) {
		if strings.Contains(flagValue, "parseTime=true") {
			dburi = flagValue
			return nil
		}
		return errors.New("please add 'parseTime=true' to the dburi")
	})
	flag.BoolVar(&initFlag, "initsql", false, "Create database schema and tables")
	flag.Parse()

	if initFlag {
		err := initDB()
		if err != nil {
			panic(err)
		}
	}

	var uiFS = fs.FS(uiFiles)
	uiContent, err := fs.Sub(uiFS, "ui/dist")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(uiContent))

	http.Handle("/", fs)
	http.HandleFunc("/api/posts", postsApi)
	http.HandleFunc("/api/posts/", postApi)
	http.HandleFunc("/api/comments/", commentsApi)

	log.Printf("Running application on http://127.0.0.1:8080 with database uri configured as %s", dburi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
