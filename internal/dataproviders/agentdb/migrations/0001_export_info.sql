-- +goose Up

CREATE TABLE export_infos (
    book_id TEXT NOT NULL,
    book_url TEXT,
    relative_path TEXT NOT NULL,
    exported_at BIGINT,
    PRIMARY KEY(book_id, relative_path)
);

CREATE TABLE missing_infos (
    relative_path TEXT NOT NULL PRIMARY KEY,
    scanned_at BIGINT,
    max_entry_percentage DOUBLE NOT NULL
);
