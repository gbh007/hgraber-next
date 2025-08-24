-- +goose Up

CREATE TABLE attribute_remaps (
    attr TEXT NOT NULL REFERENCES attributes (code) ON UPDATE CASCADE ON DELETE CASCADE,
    value TEXT NOT NULL,
    to_attr TEXT REFERENCES attributes (code) ON UPDATE CASCADE ON DELETE CASCADE,
    to_value TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    PRIMARY KEY (attr, value)
);


-- +goose Down

DROP TABLE attribute_remaps;