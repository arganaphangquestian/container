-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: GetTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: CreateTodo :one
INSERT INTO todos (
  title
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM todos
WHERE id = $1;