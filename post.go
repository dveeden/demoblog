package main

import (
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Id       uint64
	Title    string
	Summary  string
	Content  string
	Comments []Comment
	Rendered template.HTML
}

// Render renders the markdown stored in `Content` and stores it in `Rendered`
func (p *Post) Render() {
	mdParser := parser.NewWithExtensions(parser.CommonExtensions)
	htmlRenderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	doc := mdParser.Parse([]byte(p.Content))

	// template.HTML() is needed to avoid HTML Escaping this.
	p.Rendered = template.HTML(string(markdown.Render(doc, htmlRenderer)))
}
