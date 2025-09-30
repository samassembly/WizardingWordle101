-- +goose Up
CREATE TYPE school AS ENUM ('fire', 'ice', 'storm', 'life', 'myth', 'death', 'balance', 'sun', 'moon', 'star', 'shadow');
CREATE TYPE s_type AS ENUM ('damage', 'heal', 'steal', 'ward', 'charm', 'global', 'manipulation', 'enchantment');

CREATE TABLE spells(
    name TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    cost INT NOT NULL,
    spell_school school,
    accuracy INT NOT NULL,
    spell_type s_type
);

-- +goose Down
DROP TABLE spells;