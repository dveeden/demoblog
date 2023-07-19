package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "embed"

	"log"
)

//go:embed html/analytics.html
var analyticsTmpl string

func analyticsPage(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.New("analytics").Parse(analyticsTmpl))

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:4000)/blog")
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database connection failed"))
		return
	}
	defer db.Close()

	err = t.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
