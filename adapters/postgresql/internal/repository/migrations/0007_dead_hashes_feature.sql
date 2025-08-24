-- +goose Up

CREATE TABLE dead_hashes (
    md5_sum TEXT NOT NULL,
    sha256_sum TEXT NOT NULL,
    size INT8 NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (md5_sum, sha256_sum, size)
);

CREATE INDEX dead_hash_md5_sum ON dead_hashes using hash (md5_sum);

-- +goose Down

DROP INDEX dead_hash_md5_sum;

DROP TABLE dead_hashes;