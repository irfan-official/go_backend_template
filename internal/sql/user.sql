-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING id, name, email;

-- name: GetUsers :many
SELECT id, name, email FROM users;