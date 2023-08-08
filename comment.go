package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	Id      uint64
	Comment string
	PostID  uint64
}

func CommentsById(id uint64) (comments []Comment, err error) {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return comments, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,body FROM comments WHERE post_id = ? ORDER BY created DESC", id)
	if err != nil {
		return comments, err
	}

	var comment Comment
	for rows.Next() {
		err = rows.Scan(&comment.Id, &comment.Comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (c *Comment) Store() (id int64, err error) {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return
	}
	defer db.Close()

	var res sql.Result
	if c.Id == 0 {
		res, err = db.Exec("INSERT INTO comments(post_id, body) VALUES(?,?)",
			c.PostID, c.Comment)
	} else {
		res, err = db.Exec("INSERT INTO comments(id, post_id, body) VALUES(?,?,?)",
			c.Id, c.PostID, c.Comment)
	}
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}
