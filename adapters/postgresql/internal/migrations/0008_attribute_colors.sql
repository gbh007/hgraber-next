-- +goose Up

CREATE TABLE attribute_colors (
    attr TEXT NOT NULL REFERENCES attributes (code) ON UPDATE CASCADE ON DELETE CASCADE,
    value TEXT NOT NULL,
    text_color VARCHAR(10) NOT NULL,
    background_color VARCHAR(10) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (attr, value)
);


-- +goose Down

DROP TABLE attribute_colors;