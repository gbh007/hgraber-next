-- +goose Up

CREATE TABLE massload_flags (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL
);

INSERT INTO massload_flags (code, name, description, created_at) VALUES 
('deduplicated', 'Дедуплицирована', 'Признак того что массовая задача дедуплицирована', NOW()),
('loaded', 'Загружена', 'Признак того что массовая задача уже была загружена в систему', NOW()),
('to_download', 'Запланирована', 'Признак того что массовая задача планируется к загрузке в систему', NOW());

ALTER TABLE massloads ADD COLUMN flags TEXT ARRAY;

UPDATE massloads SET flags = '{"deduplicated"}' WHERE is_deduplicated;

ALTER TABLE massloads DROP COLUMN is_deduplicated;


-- +goose Down

ALTER TABLE massloads ADD COLUMN is_deduplicated BOOLEAN NOT NULL DEFAULT FALSE;

UPDATE massloads SET is_deduplicated = TRUE WHERE 'deduplicated' = ANY(flags);

ALTER TABLE massloads DROP COLUMN flags;

DROP TABLE massload_flags;