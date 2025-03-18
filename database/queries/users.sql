-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at, api_key)
VALUES ($1, $2, DEFAULT, DEFAULT, DEFAULT)
RETURNING *;
