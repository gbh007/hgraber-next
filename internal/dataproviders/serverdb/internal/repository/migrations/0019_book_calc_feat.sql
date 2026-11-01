-- +goose Up

ALTER TABLE books
ADD COLUMN calc_page_count INT,
ADD COLUMN calc_file_count INT,
ADD COLUMN calc_dead_hash_count INT,
ADD COLUMN calc_page_size INT8,
ADD COLUMN calc_file_size INT8,
ADD COLUMN calc_dead_hash_size INT8,
ADD COLUMN calculated_at TIMESTAMPTZ;


-- +goose Down

ALTER TABLE books
DROP COLUMN calc_page_count,
DROP COLUMN calc_file_count,
DROP COLUMN calc_dead_hash_count,
DROP COLUMN calc_page_size,
DROP COLUMN calc_file_size,
DROP COLUMN calc_dead_hash_size,
DROP COLUMN calculated_at;
