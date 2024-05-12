-- name: CreateUser :one
INSERT INTO "user" (
    name,
    email
) VALUES (
    $1,$2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateUserEmail :exec
UPDATE "user" 
SET email = $2
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;