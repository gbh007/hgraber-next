-- +goose Up

ALTER TABLE massload_flags
ADD COLUMN order_weight INT NOT NULL DEFAULT 0,
ADD COLUMN text_color VARCHAR(10),
ADD COLUMN background_color VARCHAR(10);


ALTER TABLE massload_external_links
ADD CONSTRAINT massload_external_links_url_unique UNIQUE(url, massload_id);


ALTER TABLE massload_attributes
ADD CONSTRAINT massload_attributes_attr_unique UNIQUE(attr_code, attr_value, massload_id);

-- +goose Down


ALTER TABLE massload_flags
DROP COLUMN order_weight,
DROP COLUMN text_color,
DROP COLUMN background_color;


ALTER TABLE massload_external_links
DROP CONSTRAINT massload_external_links_url_unique;

ALTER TABLE massload_attributes
DROP CONSTRAINT massload_attributes_attr_unique;