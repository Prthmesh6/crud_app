-- name: CreateTask :one
INSERT INTO "task" (
    name,
    description,
    status,
    priority,
    due_date
) VALUES (
    $1,$2,$3,$4,$5
) RETURNING *;

-- name: GetTask :one
SELECT * FROM "task"
WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM "task" 
ORDER BY priority
LIMIT $1
OFFSET $2;

-- name: UpdateTaskStatus :exec
UPDATE "task" 
SET status = $2
WHERE id = $1;
