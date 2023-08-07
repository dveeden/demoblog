package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	Id      uint64
	Comment string
	// Author  string
	// Created time.Time
}

func CommentsById(id uint64) (comments []Comment, err error) {
	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return comments, err
	}

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
