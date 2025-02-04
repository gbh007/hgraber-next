package model

import "github.com/jackc/pgx/v5"

var _ pgx.RowScanner = (RowScanner)(nil)

type RowScanner func(rows pgx.Rows) error

func (scanner RowScanner) ScanRow(rows pgx.Rows) error {
	return scanner(rows)
}
