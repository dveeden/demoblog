package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func postsApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	posts, err := Posts(0)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to fetch posts"))
		return
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to encode posts"))
		return
	}
}

func postApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	postId, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/api/posts/"), 10, 64)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get postid from URL"))
		return
	}

	post, err := PostById(postId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to retrieve post"))
		return
	}

	post.Render()

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to encode post"))
		return
	}
}

func commentsApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	postId, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/api/comments/"), 10, 64)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get postid from URL"))
		return
	}

	comments, err := CommentsById(postId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to retrieve post"))
		return
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to encode post"))
		return
	}
}
