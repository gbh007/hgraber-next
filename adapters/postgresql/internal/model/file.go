package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var FileTable File

type File struct{}

func (File) Name() string {
	return "files"
}

func (File) ColumnID() string          { return "id" }
func (File) ColumnFilename() string    { return "filename" }
func (File) ColumnExt() string         { return "ext" }
func (File) ColumnMd5Sum() string      { return "md5_sum" }
func (File) ColumnSha256Sum() string   { return "sha256_sum" }
func (File) ColumnSize() string        { return "\"size\"" }
func (File) ColumnFSID() string        { return "fs_id" }
func (File) ColumnInvalidData() string { return "invalid_data" }
func (File) ColumnCreateAt() string    { return "create_at" }

func (f File) Columns() []string {
	return []string{
		f.ColumnID(),
		f.ColumnFilename(),
		f.ColumnExt(),
		f.ColumnMd5Sum(),
		f.ColumnSha256Sum(),
		f.ColumnSize(),
		f.ColumnFSID(),
		f.ColumnInvalidData(),
		f.ColumnCreateAt(),
	}
}

func (File) Scanner(file *core.File) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			md5Sum    sql.NullString
			sha256Sum sql.NullString
			size      sql.NullInt64
		)

		err := rows.Scan(
			&file.ID,
			&file.Filename,
			&file.Ext,
			&md5Sum,
			&sha256Sum,
			&size,
			&file.FSID,
			&file.InvalidData,
			&file.CreateAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		file.Md5Sum = md5Sum.String
		file.Sha256Sum = sha256Sum.String
		file.Size = size.Int64

		return nil
	}
}
