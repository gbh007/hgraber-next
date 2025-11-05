package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var FileTable File

type File struct {
	rawPrefix string
	prefix    string
}

func (File) WithPrefix(p string) File {
	if p == "" {
		return File{}
	}

	return File{
		rawPrefix: p,
		prefix:    p + ".",
	}
}

func (f File) Prefix() string { return f.rawPrefix }

func (File) Name() string {
	return "files"
}

func (f File) NameAlter() string {
	if f.rawPrefix == "" || f.rawPrefix == f.Name() {
		return f.Name()
	}

	return f.Name() + " " + f.rawPrefix
}

func (f File) ColumnID() string          { return f.prefix + "id" }
func (f File) ColumnFilename() string    { return f.prefix + "filename" }
func (f File) ColumnExt() string         { return f.prefix + "ext" }
func (f File) ColumnMd5Sum() string      { return f.prefix + "md5_sum" }
func (f File) ColumnSha256Sum() string   { return f.prefix + "sha256_sum" }
func (f File) ColumnSize() string        { return f.prefix + "\"size\"" }
func (f File) ColumnFSID() string        { return f.prefix + "fs_id" }
func (f File) ColumnInvalidData() string { return f.prefix + "invalid_data" }
func (f File) ColumnCreateAt() string    { return f.prefix + "create_at" }

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
