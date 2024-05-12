// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: task.sql

package db

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :one
INSERT INTO "task" (
    name,
    description,
    status,
    priority,
    due_date
) VALUES (
    $1,$2,$3,$4,$5
) RETURNING id, name, description, status, priority, due_date
`

type CreateTaskParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Status      int32          `json:"status"`
	Priority    int32          `json:"priority"`
	DueDate     int64          `json:"due_date"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.Name,
		arg.Description,
		arg.Status,
		arg.Priority,
		arg.DueDate,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.DueDate,
	)
	return i, err
}

const getTask = `-- name: GetTask :one
SELECT id, name, description, status, priority, due_date FROM "task"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id int64) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.DueDate,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, name, description, status, priority, due_date FROM "task" 
ORDER BY priority
LIMIT $1
OFFSET $2
`

type ListTasksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTasks(ctx context.Context, arg ListTasksParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.DueDate,
		); err != nil {
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

const updateTaskStatus = `-- name: UpdateTaskStatus :exec
UPDATE "task" 
SET status = $2
WHERE id = $1
`

type UpdateTaskStatusParams struct {
	ID     int64 `json:"id"`
	Status int32 `json:"status"`
}

func (q *Queries) UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateTaskStatus, arg.ID, arg.Status)
	return err
}
