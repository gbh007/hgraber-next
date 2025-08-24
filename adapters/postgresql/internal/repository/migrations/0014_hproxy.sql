-- +goose Up

ALTER TABLE agents
ADD COLUMN has_hproxy BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down

ALTER TABLE agents
DROP COLUMN has_hproxy;
