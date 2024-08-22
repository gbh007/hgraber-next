-- +goose Up
CREATE TABLE book_origin_attributes (
    book_id UUID NOT NULL REFERENCES books (id) ON UPDATE CASCADE ON DELETE CASCADE,
    attr TEXT NOT NULL REFERENCES "attributes" (code) ON UPDATE CASCADE ON DELETE CASCADE,
    "values" TEXT ARRAY NOT NULL,
    PRIMARY KEY (book_id, attr)
);

INSERT INTO
    book_origin_attributes
SELECT book_id, attr, ARRAY_AGG(value) AS "values"
FROM book_attributes
GROUP BY
    book_id,
    attr;

CREATE TABLE deleted_pages (
    book_id UUID NOT NULL REFERENCES books (id) ON UPDATE CASCADE ON DELETE CASCADE,
    page_number INT NOT NULL,
    ext TEXT NOT NULL,
    origin_url TEXT,
    md5_sum TEXT,
    sha256_sum TEXT,
    size INT8,
    downloaded BOOL NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL,
    loaded_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (book_id, page_number)
);

DROP INDEX unparsed_books;

ALTER TABLE books
ADD COLUMN deleted BOOL NOT NULL DEFAULT false,
ADD COLUMN deleted_at TIMESTAMPTZ,
ADD COLUMN verified BOOL NOT NULL DEFAULT false,
ADD COLUMN verified_at TIMESTAMPTZ;

UPDATE books SET verified = TRUE, verified_at = create_at;

CREATE INDEX unparsed_books ON books (id)
WHERE (
        name IS NULL
        OR page_count IS NULL
        OR attributes_parsed = FALSE
    )
    AND origin_url IS NOT null
    AND deleted = FALSE;

CREATE INDEX unverified_books ON books (id) WHERE verified = FALSE;

ALTER TABLE agents
ADD COLUMN can_parse_multi BOOL NOT NULL DEFAULT false;
-- +goose Down

DROP TABLE book_origin_attributes;

DROP INDEX book_attributes_value_pairs;

DROP TABLE deleted_pages;

DROP INDEX unparsed_books;

ALTER TABLE books
DROP COLUMN deleted,
DROP COLUMN deleted_at,
DROP COLUMN verified,
DROP COLUMN verified_at;

CREATE INDEX unparsed_books ON books (id)
WHERE (
        name IS NULL
        OR page_count IS NULL
        OR attributes_parsed = FALSE
    )
    AND origin_url IS NOT NULL;

DROP INDEX unverified_books;

ALTER TABLE agents DROP COLUMN can_parse_multi;