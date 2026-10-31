-- +goose Up

ALTER TABLE books
ADD COLUMN is_rebuild BOOL NOT NULL DEFAULT false;

CREATE INDEX book_is_rebuild ON books using hash (is_rebuild) WHERE is_rebuild = TRUE;

-- +goose Down

ALTER TABLE books
DROP COLUMN is_rebuild;

DROP INDEX book_is_rebuild;