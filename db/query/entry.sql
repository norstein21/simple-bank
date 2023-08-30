-- name: CreateEntry :one
-- Description: Create a new entry
INSERT INTO entries(
    accountid,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE accountid = $1
ORDER BY id
LIMIT $2
OFFSET $3;
