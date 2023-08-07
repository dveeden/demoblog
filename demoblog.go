package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"strings"
)

var dburi string = "root@tcp(127.0.0.1:4000)/blog?parseTime=true"

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

	http.HandleFunc("/api/posts", postsApi)
	http.HandleFunc("/api/posts/", postApi)
	http.HandleFunc("/api/comments/", commentsApi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
