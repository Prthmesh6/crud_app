// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (
    name,
    email
) VALUES (
    $1,$2
) RETURNING id, name, email
`

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email FROM "user"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
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

const updateUserEmail = `-- name: UpdateUserEmail :exec
UPDATE "user" 
SET email = $2
WHERE id = $1
`

type UpdateUserEmailParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateUserEmail, arg.ID, arg.Email)
	return err
}
