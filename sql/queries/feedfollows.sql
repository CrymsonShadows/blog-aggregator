-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds
ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users
ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
WITH users_feed_follows AS (
    SELECT * FROM feed_follows
    WHERE $1 = feed_follows.user_id
)
SELECT users_feed_follows.*, feeds.name AS feed_name, users.name AS user_name
from users_feed_follows
INNER JOIN feeds
ON users_feed_follows.feed_id = feeds.id
INNER JOIN users
ON users_feed_follows.user_id = users.id;

-- name: DeleteFeedFollowWithUserAndURL :exec
DELETE FROM feed_follows
USING feeds
WHERE feeds.id = feed_follows.feed_id AND $1 = feed_follows.user_id AND $2 = feeds.url;