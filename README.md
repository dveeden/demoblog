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
    - Show the number of likes on posts
    - Show the number of likes on comments
    - Storing a cached version of the rendered markdown? Could be used for the summary on the index page.
    - TiCDC & Kafka for comments, likes, heartbeat (alert?)
    - Load generating script/tool
- Analytics
    - Number of comments per post

    ```
    SELECT
        p.title,
        COUNT(c.id)
    FROM
        posts p
    LEFT JOIN
        comments c ON p.id=c.post_id
    GROUP BY
        p.id
    ```

    - Number of comments per author
    - Most liked posts/comments

    ```
    SELECT
        p.title,
        p.likes,
        SUM(c.likes) 'comment likes'
    FROM
        posts p
    JOIN
        comments c ON c.post_id=p.id
    GROUP BY
        p.id
    ORDER BY
        p.likes+SUM(c.likes) DESC
    LIMIT 5;
    ```
    - Average length of comment, post
    - Comments per hour of day


# Things to try

- Block writes

```sql
BEGIN;
SELECT * FROM ticker FOR UPDATE;
DO SLEEP(10);
ROLLBACK;
```

This should show increase in the ticker latency. Other things like DDL etc should not show this.

- Add/remove TiFlash replicas
- Add/remove index on `comments.post_id`
