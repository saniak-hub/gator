-- name: CreatePost :one
INSERT INTO posts(
    id, created_at, updated_at, title, url, description, published_at, feed_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *
;


-- name: GetUserPosts :many
SELECT title, posts.url, description FROM posts
JOIN feeds ON feeds.id = posts.feed_id
JOIN users ON users.id = feeds.user_id
WHERE users.id = $1
ORDER BY posts.created_at
LIMIT $2
;
