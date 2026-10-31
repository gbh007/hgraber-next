-- +goose Up

CREATE INDEX file_md5_sum ON files using hash (md5_sum);
CREATE INDEX book_attributes_value ON book_attributes using hash (value);
CREATE INDEX book_labels_value ON book_labels using hash (value);

-- +goose Down

DROP INDEX file_md5_sum;
DROP INDEX book_attributes_value;
DROP INDEX book_labels_value;