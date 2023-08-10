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
	Likes    uint64
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

func (p *Post) Like() error {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE posts SET likes=likes+1 WHERE id=?", p.Id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Post) Store() (id int64, err error) {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return
	}
	defer db.Close()

	authorId := 1
	if p.Id == 0 {
		_, err = db.Exec("INSERT INTO posts(title,body,author_id) VALUES(?,?,?)",
			p.Title, p.Body, authorId)
	} else {
		_, err = db.Exec("INSERT INTO posts(id,title,body,author_id) VALUES(?,?,?)",
			p.Id, p.Title, p.Body, authorId)
	}
	return
}

func PostById(id uint64) (post Post, err error) {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return post, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id,title,body,likes FROM posts WHERE id = ?", id)

	err = row.Scan(&post.Id, &post.Title, &post.Body, &post.Likes)
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
	defer db.Close()

	var p Post
	rows, err := db.Query(
		"SELECT id,title,body,likes FROM posts WHERE id > ? ORDER BY created DESC", start,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Title, &p.Body, &p.Likes)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return
}
