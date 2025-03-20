-- name: CreateFeedFollowers :one
INSERT INTO feed_followers (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, DEFAULT, DEFAULT, $2, $3)
RETURNING *;

-- name: GetFeedFollowers :many
SELECT * FROM feed_followers WHERE user_id=$1;

-- name: DeleteFeedFollowers :exec
DELETE FROM feed_followers WHERE id=$1 AND user_id=$2;