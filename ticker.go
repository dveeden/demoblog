package main

import (
	"database/sql"
	"log"
	"time"
)

func ticker() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:4000)/blog")
	if err != nil {
		log.Fatalf("Failed to setup database connection for ticker: %s", err)
	}
	defer db.Close()

	for {
		_, err = db.Exec(
			"INSERT INTO ticker(id) VALUES(1) " +
				"ON DUPLICATE KEY UPDATE " +
				"ts=CURRENT_TIMESTAMP(6)")
		if err != nil {
			log.Print(err)
		}
		time.Sleep(time.Second * 1)
	}
}
