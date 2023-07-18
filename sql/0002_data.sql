USE blog;

INSERT INTO authors(id, name)
VALUES(1, "John Doe");

INSERT INTO posts(id, author_id, title, body) VALUES
(1, 1, "First test post", "Test post body for post 1"),
(2, 1, "Second test post", "Test post body for post 2"),
(3, 1, "MD Demo", "# Overview\nTest *for* _markdown_\n```\nfoo\n```\nThis works: [click here](https://ossinsight.io)");

INSERT INTO comments(post_id, body, author_id) VALUES
(1, 'test comment', NULL),
(1, 'Another comment', NULL),
(1, 'Third comment', 1);