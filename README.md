To run this:

Run a local TiDB database on port 40000

```
tiup playground v7.2.0
```

Create the database schema

```
mysql --comments -h 127.0.0.1 -P 4000 -u root
...
CREATE SCHEMA blog COLLATE utf8mb4_bin;
```

Now build and run the code

```
go build
./demoblog
```

And finally open the webpage in your browser.
```
xdg-open http://127.0.0.1:8080/
```

# HTTP API

| Endpoint | HTTP Verb | Description |
|----------|-----------|-------------|
| /api/posts | GET | List of posts |
| /api/posts/:postid | GET | Get rendered post |
| /api/posts/:postid/like | POST | Increase post likes |
| /api/comments/:postid | GET | Get comments for post |
| /api/comments/:postid | POST | Store comment for post |


# Development

The `./demoblog` binary holds both the backend code for the HTTP API and the files for the frontend. Just running this
binary is sufficient for normal operation. However you can run `npm run dev` in the `ui` directory to get development 
environment with automatic reloading based on [Vite](https://vitejs.dev/).
# TODO

- General
    - Editing posts
    - Pagination for the list of posts
    - Show the number of likes on posts
    - Show the number of likes on comments
    - TiCDC & Kafka?
- Load generator
    - Write one
- Analytics
    - Number of comments per author
    - Most liked posts/comments
    - Average length of comment, post
    - Comments per hour of day


# Things to try

- Block writes with a `SELECT ... FOR UPDATE`
- Add/remove TiFlash replicas
- Add/remove index on `comments.post_id`
- Try foreign keys
