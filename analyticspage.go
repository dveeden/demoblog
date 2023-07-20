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

type analytics struct {
	db              *sql.DB
	CommentsPerPost map[string]uint64
}

func (a *analytics) commentsPerPost() {
	var (
		title string
		count uint64
	)
	q :=
		`SELECT
			p.title,
			COUNT(c.id)
		FROM
			posts p
		LEFT JOIN
			comments c ON p.id=c.post_id
		GROUP BY
			p.id`

	a.CommentsPerPost = make(map[string]uint64, 10)
	rows, err := a.db.Query(q)
	for rows.Next() {
		if err != nil {
			log.Print(err)
		} else {
			rows.Scan(&title, &count)
			a.CommentsPerPost[title] = count
		}
	}
}

func analyticsPage(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.New("analytics").Parse(analyticsTmpl))

	db, err := sql.Open("mysql", dburi)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database connection failed"))
		return
	}
	defer db.Close()

	a := analytics{db: db}
	a.commentsPerPost()

	err = t.Execute(w, a)
	if err != nil {
		log.Print(err)
	}
}
