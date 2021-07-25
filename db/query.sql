CREATE TABLE users(
	id BIGSERIAL PRIMARY KEY,
	username VARCHAR(25) UNIQUE NOT NULL,
	password_hash VARCHAR(36) NOT NULL,
	first_name VARCHAR(25) NOT NULL,
	last_name VARCHAR(25) NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL
);
CREATE INDEX ON users(username);

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
	first_name, last_name, username, password_hash
) VALUES (
	$1, $2, $3, $4
)
RETURNING *;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteUserByUsername :exec
DELETE FROM users
WHERE username = $1;
