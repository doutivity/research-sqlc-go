-- name: UserNew :exec
INSERT INTO users (name, created_at)
VALUES (@name, @created_at);

-- name: UserNewAndGet :one
INSERT INTO users (name, created_at)
VALUES (@name, @created_at)
RETURNING id, name, created_at;

-- name: UserGetByID :one
SELECT id, name, created_at
FROM users
WHERE id = @id;

-- name: UsersCount :one
SELECT COUNT(*)
FROM users;

-- name: Users :many
SELECT id, name, created_at
FROM users
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
