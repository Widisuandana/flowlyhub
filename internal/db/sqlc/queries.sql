-- name: CreateUser :one
INSERT INTO users (email, password, role, name)
VALUES ($1, $2, $3, $4)
RETURNING id, email, role, name, created_at;

-- name: GetUserByEmail :one
SELECT id, email, password, role, name, created_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, email, role, name, created_at
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET email = $2, password = $3, role = $4, name = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, email, role, name, created_at, updated_at;

-- name: ListUsers :many
SELECT id, name, email, role FROM users ORDER BY id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
