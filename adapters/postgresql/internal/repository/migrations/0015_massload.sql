-- +goose Up

CREATE TABLE massloads (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    is_deduplicated BOOLEAN NOT NULL DEFAULT FALSE,
    page_size INT8,
    file_size INT8,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);

CREATE TABLE massload_external_links (
    massload_id INT NOT NULL REFERENCES massloads (id) ON UPDATE CASCADE ON DELETE CASCADE,
    url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE massload_attributes (
    massload_id INT NOT NULL REFERENCES massloads (id) ON UPDATE CASCADE ON DELETE CASCADE,
    attr_code TEXT NOT NULL REFERENCES attributes (code) ON UPDATE CASCADE ON DELETE CASCADE,
    attr_value TEXT NOT NULL,
    page_size INT8,
    file_size INT8,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);


-- +goose Down

DROP TABLE massload_attributes;
DROP TABLE massload_external_links;
DROP TABLE massloads;