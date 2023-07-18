package main

import (
	"context"
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
	db.Exec("SET time_zone='UTC'")

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

func tickerChecker() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:4000)/blog?parseTime=true")
	if err != nil {
		log.Fatalf("Failed to setup database connection for ticker checker: %s", err)
	}
	defer db.Close()
	db.Exec("SET time_zone='UTC'")

	var ts time.Time
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		row := db.QueryRowContext(ctx, "SELECT ts FROM ticker WHERE id=1")
		err = row.Scan(&ts)
		if err != nil {
			log.Printf("TickerChecker failed: %s", err)
		} else {
			log.Printf("ticker latency: %-36s is %s ago\n", ts, time.Since(ts))
		}

		time.Sleep(time.Second * 1)
	}
}
