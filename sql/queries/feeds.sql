-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;


-- name: GetFeeds :many
SELECT f.name, f.url, u.name FROM feeds f
JOIN users u
    ON u.id = f.user_id
;

-- name: GetFeedWithUrl :one
SELECT * FROM feeds
WHERE url = $1
;

-- name: UpdateFeed :exec
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW()
WHERE id = $1
;


-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at DESC
LIMIT 1;
