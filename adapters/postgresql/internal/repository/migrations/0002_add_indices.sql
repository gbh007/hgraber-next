-- +goose Up

CREATE INDEX unparsed_books ON books (id)
WHERE (
        name IS NULL
        OR page_count IS NULL
        OR attributes_parsed = FALSE
    )
    AND origin_url IS NOT NULL;

CREATE INDEX unloaded_pages ON pages (book_id, page_number)
WHERE
    downloaded = FALSE;

CREATE INDEX page_without_file ON pages (book_id, page_number)
WHERE
    file_id IS NULL;

CREATE INDEX page_with_file ON pages (file_id, book_id, page_number)
WHERE
    file_id IS NOT NULL;

CREATE INDEX page_origin_url ON pages (origin_url)
WHERE
    origin_url IS NOT NULL;

CREATE INDEX book_origin_url ON books (origin_url, id)
WHERE
    origin_url IS NOT NULL;

CREATE INDEX attribute_book_id_code ON book_attributes (book_id, attr);

CREATE INDEX file_unhandled ON files (id)
WHERE
    md5_sum IS NULL
    OR sha256_sum IS NULL
    OR "size" IS NULL;

-- +goose Down

DROP INDEX unparsed_books;

DROP INDEX unloaded_pages;

DROP INDEX page_without_file;

DROP INDEX page_with_file;

DROP INDEX page_origin_url;

DROP INDEX book_origin_url;

DROP INDEX attribute_book_id_code;

DROP INDEX file_unhandled;