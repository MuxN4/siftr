-- name: CreateFeedFollowers :one
INSERT INTO feed_followers (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, DEFAULT, DEFAULT, $2, $3)
RETURNING *;

