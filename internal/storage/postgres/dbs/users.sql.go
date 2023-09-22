// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: users.sql

package dbs

import (
	"context"
	"time"
)

const userGetByID = `-- name: UserGetByID :one
SELECT id, name, created_at
FROM users
WHERE id = $1
`

func (q *Queries) UserGetByID(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.userGetByIDStmt, userGetByID, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const userNew = `-- name: UserNew :exec
INSERT INTO users (name, created_at)
VALUES ($1, $2)
`

type UserNewParams struct {
	Name      string
	CreatedAt time.Time
}

func (q *Queries) UserNew(ctx context.Context, arg UserNewParams) error {
	_, err := q.exec(ctx, q.userNewStmt, userNew, arg.Name, arg.CreatedAt)
	return err
}

const userNewAndGet = `-- name: UserNewAndGet :one
INSERT INTO users (name, created_at)
VALUES ($1, $2)
RETURNING id, name, created_at
`

type UserNewAndGetParams struct {
	Name      string
	CreatedAt time.Time
}

func (q *Queries) UserNewAndGet(ctx context.Context, arg UserNewAndGetParams) (User, error) {
	row := q.queryRow(ctx, q.userNewAndGetStmt, userNewAndGet, arg.Name, arg.CreatedAt)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const users = `-- name: Users :many
SELECT id, name, created_at
FROM users
LIMIT $2 OFFSET $1
`

type UsersParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) Users(ctx context.Context, arg UsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.usersStmt, users, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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

const usersCount = `-- name: UsersCount :one
SELECT COUNT(*)
FROM users
`

func (q *Queries) UsersCount(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.usersCountStmt, usersCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}