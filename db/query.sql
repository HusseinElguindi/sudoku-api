CREATE TABLE users(
	id BIGSERIAL PRIMARY KEY,
	username VARCHAR(25) UNIQUE NOT NULL,
	password_hash VARCHAR(36) NOT NULL,
	first_name VARCHAR(25) NOT NULL,
	last_name VARCHAR(25) NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL
);
CREATE INDEX ON users(username);

CREATE TABLE puzzles(
	id BIGSERIAL PRIMARY KEY,
	array TEXT UNIQUE NOT NULL
);
CREATE INDEX ON puzzles(array);

CREATE TABLE user_puzzles(
	user_id BIGSERIAL NOT NULL,
	puzzle_id BIGSERIAL NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(puzzle_id) REFERENCES puzzles(id),
	UNIQUE(user_id, puzzle_id)
);
CREATE INDEX ON user_puzzles(user_id);
CREATE INDEX ON user_puzzles(puzzle_id);
CREATE INDEX ON user_puzzles(user_id, puzzle_id);

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
