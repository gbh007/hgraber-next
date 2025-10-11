package model

import (
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var DeadHashTable DeadHash

type DeadHash struct{}

func (DeadHash) Name() string {
	return "dead_hashes"
}

func (DeadHash) ColumnMd5Sum() string    { return "md5_sum" }
func (DeadHash) ColumnSha256Sum() string { return "sha256_sum" }
func (DeadHash) ColumnSize() string      { return "size" }
func (DeadHash) ColumnCreatedAt() string { return "created_at" }

func (a DeadHash) Columns() []string {
	return []string{
		a.ColumnMd5Sum(),
		a.ColumnSha256Sum(),
		a.ColumnSize(),
		a.ColumnCreatedAt(),
	}
}

func (DeadHash) Scanner(hash *core.DeadHash) RowScanner {
	return func(rows pgx.Rows) error {
		err := rows.Scan(
			&hash.Md5Sum,
			&hash.Sha256Sum,
			&hash.Size,
			&hash.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		return nil
	}
}
