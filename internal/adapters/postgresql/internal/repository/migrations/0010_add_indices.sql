-- +goose Up

DROP INDEX page_origin_url;
DROP INDEX book_origin_url;
DROP INDEX book_is_rebuild;



CREATE INDEX page_origin_url ON pages using hash (origin_url)
WHERE origin_url IS NOT NULL;

CREATE INDEX book_origin_url ON books using hash (origin_url)
WHERE origin_url IS NOT NULL;

CREATE INDEX book_is_rebuild ON books (id)
WHERE is_rebuild = true;


CREATE INDEX book_is_deleted ON books (id)
WHERE deleted = true;


CREATE INDEX file_invalid_data ON files (id, fs_id)
WHERE invalid_data = true;


CREATE INDEX page_only_file_id ON pages (file_id)
WHERE file_id IS NOT NULL;


-- +goose Down


DROP INDEX page_origin_url;
DROP INDEX book_origin_url;
DROP INDEX book_is_rebuild;
DROP INDEX book_is_deleted;
DROP INDEX file_invalid_data;
DROP INDEX page_only_file_id;



CREATE INDEX page_origin_url ON pages (origin_url)
WHERE
    origin_url IS NOT NULL;

CREATE INDEX book_origin_url ON books (origin_url, id)
WHERE
    origin_url IS NOT NULL;

CREATE INDEX book_is_rebuild ON books using hash (is_rebuild) WHERE is_rebuild = TRUE;