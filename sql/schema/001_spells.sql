-- +goose Up
CREATE TABLE spells(
    name TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    cost INT NOT NULL,
    spell_school TEXT NOT NULL,
    accuracy INT NOT NULL,
    spell_type TEXT NOT NULL
);

-- +goose Down
DROP TABLE spells;