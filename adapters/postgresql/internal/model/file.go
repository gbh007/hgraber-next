package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func FileColumns() []string {
	return []string{
		"id",
		"filename",
		"ext",
		"md5_sum",
		"sha256_sum",
		"\"size\"",
		"fs_id",
		"invalid_data",
		"create_at",
	}
}

func FileScanner(file *core.File) RowScanner {
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
