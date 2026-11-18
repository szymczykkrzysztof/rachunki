-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING id, email, password_hash, role, created_at;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, role, created_at
FROM users
WHERE lower(email) = lower(sqlc.arg(email));

-- name: GetUserByID :one
SELECT id, email, password_hash, role, created_at
FROM users
WHERE id = $1;
