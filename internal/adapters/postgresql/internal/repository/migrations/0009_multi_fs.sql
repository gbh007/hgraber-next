-- +goose Up

CREATE TABLE file_storages (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    agent_id UUID REFERENCES agents (id) ON UPDATE CASCADE ON DELETE SET NULL,
    path TEXT,
    download_priority INT NOT NULL DEFAULT 0,
    deduplicate_priority INT NOT NULL DEFAULT 0,
    highway_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    highway_addr TEXT,
    created_at TIMESTAMPTZ NOT NULL
);

ALTER TABLE files
ADD COLUMN fs_id UUID REFERENCES file_storages (id) ON UPDATE CASCADE ON DELETE CASCADE,
ADD COLUMN invalid_data BOOLEAN NOT NULL DEFAULT FALSE;

ALTER TABLE agents
ADD COLUMN has_fs BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down

ALTER TABLE agents
DROP COLUMN has_fs;

ALTER TABLE files
DROP COLUMN fs_id,
DROP COLUMN invalid_data;

DROP TABLE file_storages;