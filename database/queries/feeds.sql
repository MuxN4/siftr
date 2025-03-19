-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
VALUES ($1, $2, $3, DEFAULT, DEFAULT, $4)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

