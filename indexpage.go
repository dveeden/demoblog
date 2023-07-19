package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"

	_ "embed"

	"log"
)

//go:embed html/index.html
var indexTmpl string

func indexPage(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.New("index").Parse(indexTmpl))

	var indexContent struct {
		Dbver string
		Tick  time.Time
		Posts []Post
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:4000)/blog?parseTime=true")
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database connection failed"))
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["title"][0] != "" && r.Form["body"][0] != "" {
			_, err = db.Exec(`INSERT INTO posts(title, body, author_id) VALUES (?,?,?)`,
				r.Form["title"][0], r.Form["body"][0], 1)
			if err != nil {
				log.Print(err)
			}
		}
	}

	if err := db.QueryRow("SELECT TIDB_VERSION(), ts FROM ticker WHERE id=1").Scan(
		&indexContent.Dbver, &indexContent.Tick); err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database query for db version and ticker failed"))
		return
	}

	rows, err := db.Query("SELECT id,title,body FROM posts LIMIT 10")
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database query for posts failed"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.Id, &p.Title, &p.Content); err != nil {
			log.Print(err)
		} else {
			if len(p.Content) > 100 {
				p.Summary = p.Content[:100] + "..."
			} else {
				p.Summary = p.Content
			}
			indexContent.Posts = append(indexContent.Posts, p)
		}
	}

	err = t.Execute(w, indexContent)
	if err != nil {
		log.Print(err)
	}
}
