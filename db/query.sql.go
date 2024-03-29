// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
)

const createPuzzle = `-- name: CreatePuzzle :one
INSERT INTO puzzles (
	array_str
) VALUES (
	$1
)
RETURNING id, array_str, created_at
`

func (q *Queries) CreatePuzzle(ctx context.Context, arrayStr string) (Puzzle, error) {
	row := q.db.QueryRowContext(ctx, createPuzzle, arrayStr)
	var i Puzzle
	err := row.Scan(&i.ID, &i.ArrayStr, &i.CreatedAt)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	first_name, last_name, username, password_hash
) VALUES (
	$1, $2, $3, $4
)
RETURNING id, username, password_hash, first_name, last_name, created_at
`

type CreateUserParams struct {
	FirstName    string
	LastName     string
	Username     string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.PasswordHash,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const createUserPuzzle = `-- name: CreateUserPuzzle :one
INSERT INTO user_puzzles (
	user_id, puzzle_id
) VALUES (
	$1, $2
)
RETURNING user_id, puzzle_id, created_at
`

type CreateUserPuzzleParams struct {
	UserID   int64
	PuzzleID int64
}

func (q *Queries) CreateUserPuzzle(ctx context.Context, arg CreateUserPuzzleParams) (UserPuzzle, error) {
	row := q.db.QueryRowContext(ctx, createUserPuzzle, arg.UserID, arg.PuzzleID)
	var i UserPuzzle
	err := row.Scan(&i.UserID, &i.PuzzleID, &i.CreatedAt)
	return i, err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserByID, id)
	return err
}

const deleteUserByUsername = `-- name: DeleteUserByUsername :exec
DELETE FROM users
WHERE username = $1
`

func (q *Queries) DeleteUserByUsername(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUserByUsername, username)
	return err
}

const deleteUserPuzzle = `-- name: DeleteUserPuzzle :exec
DELETE FROM user_puzzles
WHERE user_id = $1 AND puzzle_id = $2
`

type DeleteUserPuzzleParams struct {
	UserID   int64
	PuzzleID int64
}

func (q *Queries) DeleteUserPuzzle(ctx context.Context, arg DeleteUserPuzzleParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserPuzzle, arg.UserID, arg.PuzzleID)
	return err
}

const getPuzzleByArrayStr = `-- name: GetPuzzleByArrayStr :one
SELECT id, array_str, created_at FROM puzzles
WHERE array_str = $1 LIMIT 1
`

func (q *Queries) GetPuzzleByArrayStr(ctx context.Context, arrayStr string) (Puzzle, error) {
	row := q.db.QueryRowContext(ctx, getPuzzleByArrayStr, arrayStr)
	var i Puzzle
	err := row.Scan(&i.ID, &i.ArrayStr, &i.CreatedAt)
	return i, err
}

const getPuzzleByID = `-- name: GetPuzzleByID :one
SELECT id, array_str, created_at FROM puzzles
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPuzzleByID(ctx context.Context, id int64) (Puzzle, error) {
	row := q.db.QueryRowContext(ctx, getPuzzleByID, id)
	var i Puzzle
	err := row.Scan(&i.ID, &i.ArrayStr, &i.CreatedAt)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password_hash, first_name, last_name, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password_hash, first_name, last_name, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const getUserPuzzle = `-- name: GetUserPuzzle :one
SELECT user_id, puzzle_id, created_at FROM user_puzzles
WHERE user_id = $1 AND puzzle_id = $2
LIMIT 1
`

type GetUserPuzzleParams struct {
	UserID   int64
	PuzzleID int64
}

func (q *Queries) GetUserPuzzle(ctx context.Context, arg GetUserPuzzleParams) (UserPuzzle, error) {
	row := q.db.QueryRowContext(ctx, getUserPuzzle, arg.UserID, arg.PuzzleID)
	var i UserPuzzle
	err := row.Scan(&i.UserID, &i.PuzzleID, &i.CreatedAt)
	return i, err
}

const listUserPuzzles = `-- name: ListUserPuzzles :many
SELECT user_id, puzzle_id, created_at FROM user_puzzles
WHERE user_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListUserPuzzles(ctx context.Context, userID int64) ([]UserPuzzle, error) {
	rows, err := q.db.QueryContext(ctx, listUserPuzzles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserPuzzle
	for rows.Next() {
		var i UserPuzzle
		if err := rows.Scan(&i.UserID, &i.PuzzleID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
