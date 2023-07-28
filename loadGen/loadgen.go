package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type loadGen struct {
	urlBase url.URL
	maxId   map[string]uint64
}

func (lg *loadGen) basicRequests() {

	endpoints := []struct {
		method string
		path   string
		values url.Values
		nr     int // Number of requests to this endpoint
	}{
		{"GET", "/", nil, 10},
		{"GET", "/analytics", nil, 2},
		{"GET", "/posts", url.Values{"id": []string{"1"}}, 5},
		{"GET", "/posts", url.Values{"id": []string{"aaa"}}, 1},
		{"GET", "/posts", url.Values{"id": []string{strconv.FormatUint(math.MaxUint64, 10)}}, 1},
	}

	for _, e := range endpoints {
		var err error
		var r *http.Response

		u := lg.urlBase.JoinPath(e.path)
		u.RawQuery = e.values.Encode()

		for i := 0; i < e.nr; i++ {
			switch e.method {
			case "GET":
				r, err = http.Get(u.String())
			default:
				err = errors.New("unknown HTTP method")
			}
			if err != nil {
				panic(err)
			}
			log.Printf("BASIC %s on %s %s\n", r.Status, r.Request.Method, u)
		}
	}
}

// likePosts tries to like all posts until starting by the previously discovered maxId (or 1) and
// increasing the id of the post util it doesn't get a HTTP 200 OK.
func (lg *loadGen) like(likeType string, random int) {
	u := lg.urlBase.JoinPath("/likes")
	vals := url.Values{}
	vals.Set("type", likeType)

	if random > 0 {
		for i := 0; i < random; i++ {
			postID := uint64(rand.Int63n(int64(lg.maxId[likeType]))) + 1
			vals.Set("id", strconv.FormatUint(postID, 10))
			u.RawQuery = vals.Encode()
			r, err := http.Post(u.String(), "application/json", nil)
			if err != nil {
				panic(err)
			}
			log.Printf("LIKE RANDOM: %s %s\n", r.Status, u.String())
			if r.StatusCode != 200 {
				return
			}
		}
	} else {
		postID := lg.maxId[likeType]

		for {
			vals.Set("id", strconv.FormatUint(postID, 10))
			u.RawQuery = vals.Encode()
			r, err := http.Post(u.String(), "application/json", nil)
			if err != nil {
				panic(err)
			}
			log.Printf("LIKE ORDERED: %s %s\n", r.Status, u.String())
			if r.StatusCode != 200 {
				return
			}
			if postID > lg.maxId[likeType] {
				lg.maxId[likeType] = postID
			}
			postID++
		}
	}
}

func (lg *loadGen) comment(nr int) {
	for i := 0; i < nr; i++ {
		postID := uint64(rand.Int63n(int64(lg.maxId["post"]))) + 1
		vals := url.Values{}
		vals.Set("id", strconv.FormatUint(postID, 10))
		u := lg.urlBase.JoinPath("/posts")
		u.RawQuery = vals.Encode()
		comment := url.Values{}
		comment.Set("comment", fmt.Sprintf("testcomment %d, Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", i))
		commentReader := strings.NewReader(comment.Encode())
		r, err := http.Post(u.String(), "application/x-www-form-urlencoded", commentReader)
		if err != nil {
			panic(err)
		}
		log.Printf("COMMENT: %s %s\n", r.Status, u.String())
	}
}

func main() {
	lg := loadGen{
		urlBase: url.URL{
			Scheme: "http",
			Host:   "127.0.0.1:8080",
		},
		maxId: map[string]uint64{
			"comment": 1,
			"post":    1,
		},
	}

	for {
		lg.basicRequests()

		// Unordered likes are needed to discover the max post and comment ID's
		// Running this each round to discover new ID's
		lg.like("post", 0)
		lg.like("comment", 0)

		lg.like("post", 5)
		lg.like("comment", 10)
		lg.comment(3)
		time.Sleep(time.Second * 2)
	}
}
