package model

import (
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var DeadHashTable = DeadHash{baseTable: baseTable{name: "dead_hashes"}}

type DeadHash struct {
	baseTable
}

func (dh DeadHash) WithPrefix(pf string) DeadHash {
	return DeadHash{
		baseTable: dh.withPrefix(pf),
	}
}

func (dh DeadHash) ColumnMd5Sum() string    { return dh.prefix + "md5_sum" }
func (dh DeadHash) ColumnSha256Sum() string { return dh.prefix + "sha256_sum" }
func (dh DeadHash) ColumnSize() string      { return dh.prefix + "size" }
func (dh DeadHash) ColumnCreatedAt() string { return dh.prefix + "created_at" }

func (dh DeadHash) Columns() []string {
	return []string{
		dh.ColumnMd5Sum(),
		dh.ColumnSha256Sum(),
		dh.ColumnSize(),
		dh.ColumnCreatedAt(),
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
