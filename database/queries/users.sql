-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at, api_key)
VALUES ($1, $2, DEFAULT, DEFAULT, DEFAULT)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;