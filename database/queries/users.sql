-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES ($1, $2, DEFAULT, DEFAULT)
RETURNING *;
