package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	_ "embed"

	"log"
)

//go:embed html/post.html
var postTmpl string

func postPage(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.New("post").Parse(postTmpl))
	var p Post

	idRaw := r.URL.Query().Get("id")
	if idRaw == "" {
		log.Print("Missing id")
		w.WriteHeader(404)
		w.Write([]byte("Missing id"))
		return
	}
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Conversion of id failed"))
		return
	}

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:4000)/blog")
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database connection failed"))
		return
	}
	defer db.Close()

	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["comment"][0] != "" {
			_, err = db.Exec(`INSERT INTO comments(body, post_id, author_id) VALUES (?,?,?)`,
				r.Form["comment"][0], id, nil)
			if err != nil {
				log.Print(err)
			}
		}
	}

	row := db.QueryRow("SELECT title,body FROM posts WHERE id=?", id)
	err = row.Scan(&p.Title, &p.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			w.Write([]byte("No post found with this id"))
		} else {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Database query failed"))
		}
		return
	}

	var commentJson []byte
	row = db.QueryRow(
		`SELECT
			JSON_ARRAYAGG(
					JSON_OBJECT(
						"comment", c.body,
						"created", DATE_FORMAT(created, "%Y-%m-%dT%H%:%i:%SZ"),
						"author", IFNULL(a.name, "Anonymous user")
					)
			) 
			FROM
				comments c
				LEFT JOIN authors a ON c.author_id=a.id 
			WHERE
				post_id=?`,
		id)
	err = row.Scan(&commentJson)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Database query for comments failed"))
		}
		return
	}
	if commentJson != nil {
		err = json.Unmarshal(commentJson, &p.Comments)
		if err != nil {
			log.Print(err)
		}
	}

	p.Render() // Markdown to HTML
	err = t.Execute(w, p)
	if err != nil {
		log.Print(err)
	}
}
