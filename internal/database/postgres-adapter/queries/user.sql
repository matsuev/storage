-- name: GetUsers :many
SELECT * FROM public.user
;

-- name: GetUserByID :one
SELECT * FROM public.user
WHERE id=$1
LIMIT 1
;

-- name: DeleteUserByID :exec
DELETE FROM public.user
WHERE id=$1
;