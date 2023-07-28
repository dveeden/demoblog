package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type postApiResponse struct {
	PostIds []uint64 `json:"postids"`
}

func postApi(w http.ResponseWriter, r *http.Request) {
	resp := postApiResponse{}

	db, err := sql.Open("mysql", dburi)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database connection failed"))
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id FROM posts`)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("failed to encode response"))
		return
	}
	for rows.Next() {
		var postId uint64
		rows.Scan(&postId)
		resp.PostIds = append(resp.PostIds, postId)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("failed to encode response"))
		return
	}
}
