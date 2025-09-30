-- name: CreateSpell :one
INSERT INTO spells (name, created_at, updated_at, cost, spell_school, accuracy, spell_type)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetSpell :one
SELECT * FROM spells
WHERE name = $1;

-- name: GetSpells :many
SELECT * FROM spells;

-- name: DeleteSpells :exec
DELETE FROM spells;