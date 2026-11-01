-- +goose Up

ALTER TABLE export_infos
RENAME COLUMN exported_at TO imported_at;

ALTER TABLE export_infos
RENAME TO import_infos;


-- +goose Down

ALTER TABLE import_infos
RENAME TO export_infos;

ALTER TABLE export_infos
RENAME COLUMN imported_at TO exported_at;