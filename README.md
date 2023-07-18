To run this:

Run a local TiDB database on port 40000

```
tiup playground v7.2.0
```

Create the database schema and initial data

```
mysql --comments -h 127.0.0.1 -P 4000 -u root
...
source 0001_schema.sql
source 0002_data.sql
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

# TODO

- General
    - Likes for posts
    - Likes for comments
    - Storing a cached version of the rendered markdown? Could be used for the summary on the index page.
    - TiCDC & Kafka for comments, likes, heartbeat (alert?)
    - Load generating script/tool
- Analytics
    - Number of comments per post

    ```
    SELECT p.title, COUNT(c.id) FROM posts p LEFT JOIN comments c ON p.id=c.post_id GROUP BY p.id
    ```

    - Number of comments per author

    - Likes per comment/post

    - Average length of comment, post

    - Comments per hour of day


