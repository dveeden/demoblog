package main

import (
	"database/sql"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Id       uint64
	Title    string
	Body     string
	Rendered template.HTML
}

// Render renders the markdown stored in `Content` and stores it in `Rendered`
func (p *Post) Render() {
	mdParser := parser.NewWithExtensions(parser.CommonExtensions)
	htmlRenderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	doc := mdParser.Parse([]byte(p.Body))

	// template.HTML() is needed to avoid HTML Escaping this.
	p.Rendered = template.HTML(string(markdown.Render(doc, htmlRenderer)))
}

func PostById(id uint64) (post Post, err error) {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return post, err
	}

	row := db.QueryRow("SELECT id,title,body FROM posts WHERE id = ?", id)

	err = row.Scan(&post.Id, &post.Title, &post.Body)
	if err != nil {
		return post, err
	}
	return post, nil
}

func Posts(start uint64) (posts []Post, err error) {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return nil, err
	}

	var p Post
	rows, err := db.Query("SELECT id,title,body FROM posts WHERE id > ? ORDER BY created DESC", start)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Title, &p.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return
}
