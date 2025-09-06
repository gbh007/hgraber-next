-- +goose Up

ALTER TABLE massloads
ADD COLUMN page_count INT8,
ADD COLUMN file_count INT8,
ADD COLUMN books_ahead INT8,
ADD COLUMN new_books INT8,
ADD COLUMN existing_books INT8,
ADD COLUMN books_in_system INT8;

ALTER TABLE massload_external_links
ADD COLUMN books_ahead INT8,
ADD COLUMN new_books INT8,
ADD COLUMN existing_books INT8,
ADD COLUMN auto_check BOOLEAN NOT NULL DEFAULT FALSE,
ADD COLUMN updated_at TIMESTAMPTZ;

ALTER TABLE massload_attributes
ADD COLUMN page_count INT8,
ADD COLUMN file_count INT8,
ADD COLUMN books_in_system INT8;

-- +goose Down

ALTER TABLE massloads
DROP COLUMN page_count,
DROP COLUMN file_count,
DROP COLUMN books_ahead,
DROP COLUMN new_books,
DROP COLUMN existing_books,
DROP COLUMN books_in_system;

ALTER TABLE massload_external_links
DROP COLUMN books_ahead,
DROP COLUMN new_books,
DROP COLUMN existing_books,
DROP COLUMN auto_check,
DROP COLUMN updated_at;

ALTER TABLE massload_attributes
DROP COLUMN page_count,
DROP COLUMN file_count,
DROP COLUMN books_in_system;