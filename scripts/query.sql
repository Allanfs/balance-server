-- name: GetEntries :many
SELECT * from balance.entry e;

-- name: GetEntry :one
SELECT * from balance.entry e where e.id = $1;

-- name: GetEntriesByType :many
SELECT * from balance.entry e where e.entry_type = $1;

-- name: GetEntryByExternalInfo :many
SELECT * from balance.entry e where external_info = $1;

-- name: GetEntryByDate :many
SELECT * from balance.entry e where created_at = $1;

-- name: CreateEntry :exec
INSERT INTO balance.entry (name, amount, entry_type, external_info) VALUES ($1, $2, $3, $4);

-- name: UpdateEntryName :exec
UPDATE balance.entry e SET name = $2 WHERE id = $1;

-- name: UpdateEntryAmount :exec
UPDATE balance.entry e SET amount = $2 WHERE id = $1;

UPDATE balance.entry e SET entry_type = $2 WHERE id = $1;

-- name: UpdateEntryExternalInfo :exec
UPDATE balance.entry e SET external_info = $2 WHERE id = $1;

-- name: UpdateEntry :exec
UPDATE balance.entry e SET name = $2, amount = $3, entry_type = $4, external_info = $5 WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM balance.entry e WHERE id = $1;

-- name: DeleteEntryByExternalInfo :exec
DELETE FROM balance.entry e WHERE external_info = $1;

