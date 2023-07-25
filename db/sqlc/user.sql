-- name: CreateUser :one
INSERT INTO users (
        username,
        email,
        password,
        access
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;
-- name: GetUserForUpdate :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY username;
-- name: UpdateUser :one
UPDATE users
set email = $2,
    password = $3,
    access = $4
WHERE username = $1
RETURNING *;