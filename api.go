package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var isPostLike = regexp.MustCompile(`^/api/posts/(\d*)/like`)

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

	if isPostLike.MatchString(r.URL.String()) {
		postsLikeApi(w, r)
		return
	}

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

	switch r.Method {
	case "GET":
		commentsApiGet(w, r)
	case "POST":
		commentsApiPost(w, r)
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Request method not supported"))
	}
}

func commentsApiGet(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte("Failed to encode comment"))
		return
	}
}

func commentsApiPost(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/api/comments/"), 10, 64)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get postid from URL"))
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse post data"))
		return
	}
	var comment Comment
	comment.PostID = postId
	comment.Comment = r.FormValue("Comment")
	commentId, err := comment.Store()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to store comment"))
		return
	}
	comment.Id = uint64(commentId)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to encode comment"))
		return
	}
}

func postsLikeApi(w http.ResponseWriter, r *http.Request) {
	matches := isPostLike.FindStringSubmatch(r.URL.String())
	postId, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get postid from URL"))
		return
	}

	p, err := PostById(postId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to find post with that id"))
		return
	}

	err = p.Like()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed store like"))
		return
	}
}
