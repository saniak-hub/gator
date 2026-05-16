-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feed_follow(id, created_at, updated_at, user_id, feed_id)
    VALUES($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT inserted_feed_follow.*, 
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users 
    ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds 
    ON feeds.id = inserted_feed_follow.feed_id
;

-- name: GetFeedFollowsForUser :many
SELECT 
    ff.*,
    f.name feed_name, u.name user_name
FROM feed_follow ff
INNER JOIN feeds f ON f.id = ff.feed_id
INNER JOIN users u ON u.id = ff.user_id
WHERE ff.user_id = $1
;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow
WHERE user_id = $1 
    AND feed_id = $2
;
