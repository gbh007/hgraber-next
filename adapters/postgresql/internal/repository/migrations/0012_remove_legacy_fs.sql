-- +goose Up

ALTER TABLE files
ALTER COLUMN fs_id SET NOT NULL;

-- +goose Down

ALTER TABLE files
ALTER COLUMN fs_id DROP NOT NULL;
