-- +goose Up

CREATE TABLE label_presets (
    name TEXT NOT NULL PRIMARY KEY,
    description TEXT,
    values TEXT ARRAY NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);


-- +goose Down

DROP TABLE label_presets;