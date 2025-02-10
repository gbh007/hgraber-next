-- +goose Up

CREATE TABLE url_mirrors (
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    prefixes TEXT ARRAY NOT NULL,
    description TEXT
);

-- +goose Down

DROP TABLE url_mirrors;
