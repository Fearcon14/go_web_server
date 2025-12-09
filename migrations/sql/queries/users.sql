-- name: CreateUser :one
Insert Into users (id, created_at, updated_at, email)
Values (
	gen_random_uuid(),
	now(),
	now(),
	$1
)
RETURNING *;
