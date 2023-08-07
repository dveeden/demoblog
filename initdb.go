package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() error {
	log.Println("Initializing database")

	db, err := sql.Open("mysql", dburi)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
	     id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	     created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
	     updated TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
	     author_id BIGINT UNSIGNED NOT NULL,
	     title VARCHAR(255) NOT NULL,
	     live_after TIMESTAMP(6) NULL COMMENT 'Only show post after this go-live date if set',
	     likes BIGINT UNSIGNED NOT NULL DEFAULT 0,
	     body TEXT
	
	)`)
	if err != nil {
		return fmt.Errorf("failed to create posts table: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS authors (
	     id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	     name VARCHAR(255)
	)`)
	if err != nil {
		return fmt.Errorf("failed to create authors table: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments (
	     id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	     post_id BIGINT UNSIGNED NOT NULL,
	     author_id BIGINT UNSIGNED NULL COMMENT 'Set to NULL if anonymous',
	     created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
	     likes BIGINT UNSIGNED NOT NULL DEFAULT 0,
	     body TEXT,
	     KEY(post_id)
	)`)
	if err != nil {
		return fmt.Errorf("failed to create comments table: %w", err)
	}

	// TiFlash ColumnStore
	_, err = db.Exec(`/*T! ALTER TABLE posts SET TIFLASH REPLICA 1 */`)
	if err != nil {
		log.Printf("failed to enable TiFlash replica for posts: %s", err)
	}
	_, err = db.Exec(`/*T! ALTER TABLE comments SET TIFLASH REPLICA 1 */`)
	if err != nil {
		log.Printf("failed to enable TiFlash replica for comments: %s", err)
	}

	_, err = db.Exec(`INSERT INTO authors(id, name) VALUES(1, "John Doe")`)
	if err != nil {
		log.Printf("failed to load authors: %s", err)
	}

	_, err = db.Exec(`INSERT INTO posts(id, author_id, title, body) VALUES
	(1, 1, "First test post", "Test post body for post 1"),
	(2, 1, "Second test post", "Test post body for post 2"),
	(3, 1, "MD Demo", "# Overview\nTest *for* _markdown_\n` + "```\nfoo\n```" + `\nThis works: [click here](https://ossinsight.io)")`)
	if err != nil {
		log.Printf("failed to load posts: %s", err)
	}

	_, err = db.Exec(`INSERT INTO comments(post_id, body, author_id) VALUES
	(1, 'test comment', NULL),
	(1, 'Another comment', NULL),
	(1, 'Third comment', 1)`)
	if err != nil {
		log.Printf("failed to load comments: %s", err)
	}

	return nil
}
