-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsByUserID :many
WITH users_feed_follows AS (
    SELECT * FROM feed_follows
    WHERE $1 = feed_follows.user_id
)
SELECT * FROM posts
WHERE users_feed_follows.feed_id = feed_id
ORDER BY published_at DESC LIMIT $2;