-- +goose Up

ALTER TABLE books
ADD COLUMN calc_avg_page_size INT8;


-- +goose Down

ALTER TABLE books
DROP COLUMN calc_avg_page_size;
