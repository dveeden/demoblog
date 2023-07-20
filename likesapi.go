package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

func likesApi(w http.ResponseWriter, r *http.Request) {
	likeType := r.URL.Query().Get("type")
	idRaw := r.URL.Query().Get("id")
	if idRaw == "" {
		log.Print("Missing id")
		w.WriteHeader(404)
		w.Write([]byte("Missing id"))
		return
	}
	likeId, err := strconv.Atoi(idRaw)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Conversion of id failed"))
		return
	}

	db, err := sql.Open("mysql", dburi)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		w.Write([]byte("Database connection failed"))
		return
	}
	defer db.Close()

	switch likeType {
	case "post":
		_, err = db.Exec("UPDATE posts SET likes=likes+1 WHERE id=?", likeId)
		if err != nil {
			log.Print(err)
		}
	case "comment":
		_, err = db.Exec("UPDATE comments SET likes=likes+1 WHERE id=?", likeId)
		if err != nil {
			log.Print(err)
		}
	default:
		log.Print("unknown like type")
	}

	w.WriteHeader(200)
	w.Write([]byte("liked"))
}
