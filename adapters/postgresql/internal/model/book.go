package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var BookTable Book

type Book struct{}

func (Book) Name() string {
	return "books"
}

func (Book) ColumnID() string                { return "id" }
func (Book) ColumnName() string              { return "name" }
func (Book) ColumnOriginURL() string         { return "origin_url" }
func (Book) ColumnPageCount() string         { return "page_count" }
func (Book) ColumnAttributesParsed() string  { return "attributes_parsed" }
func (Book) ColumnCreateAt() string          { return "create_at" }
func (Book) ColumnDeleted() string           { return "deleted" }
func (Book) ColumnDeletedAt() string         { return "deleted_at" }
func (Book) ColumnVerified() string          { return "verified" }
func (Book) ColumnVerifiedAt() string        { return "verified_at" }
func (Book) ColumnIsRebuild() string         { return "is_rebuild" }
func (Book) ColumnCalcPageCount() string     { return "calc_page_count" }
func (Book) ColumnCalcFileCount() string     { return "calc_file_count" }
func (Book) ColumnCalcDeadHashCount() string { return "calc_dead_hash_count" }
func (Book) ColumnCalcPageSize() string      { return "calc_page_size" }
func (Book) ColumnCalcFileSize() string      { return "calc_file_size" }
func (Book) ColumnCalcDeadHashSize() string  { return "calc_dead_hash_size" }
func (Book) ColumnCalculatedAt() string      { return "calculated_at" }
func (Book) ColumnCalcAvgPageSize() string   { return "calc_avg_page_size" }

func (b Book) Columns() []string {
	return []string{
		b.ColumnID(),
		b.ColumnName(),
		b.ColumnOriginURL(),
		b.ColumnPageCount(),
		b.ColumnAttributesParsed(),
		b.ColumnCreateAt(),
		b.ColumnDeleted(),
		b.ColumnDeletedAt(),
		b.ColumnVerified(),
		b.ColumnVerifiedAt(),
		b.ColumnIsRebuild(),
		b.ColumnCalcPageCount(),
		b.ColumnCalcFileCount(),
		b.ColumnCalcDeadHashCount(),
		b.ColumnCalcPageSize(),
		b.ColumnCalcFileSize(),
		b.ColumnCalcDeadHashSize(),
		b.ColumnCalculatedAt(),
		b.ColumnCalcAvgPageSize(),
	}
}

func (Book) Scanner(book *core.Book) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			name       sql.NullString
			originURL  sql.NullString
			pageCount  sql.NullInt32
			deletedAt  sql.NullTime
			verifiedAt sql.NullTime

			calcPageCount     sql.NullInt64
			calcFileCount     sql.NullInt64
			calcDeadHashCount sql.NullInt64
			calcPageSize      sql.NullInt64
			calcFileSize      sql.NullInt64
			calcDeadHashSize  sql.NullInt64
			calculatedAt      sql.NullTime
			calcAvgPageSize   sql.NullInt64
		)

		err := rows.Scan(
			&book.ID,
			&name,
			&originURL,
			&pageCount,
			&book.AttributesParsed,
			&book.CreateAt,
			&book.Deleted,
			&deletedAt,
			&book.Verified,
			&verifiedAt,
			&book.IsRebuild,

			&calcPageCount,
			&calcFileCount,
			&calcDeadHashCount,
			&calcPageSize,
			&calcFileSize,
			&calcDeadHashSize,
			&calculatedAt,
			&calcAvgPageSize,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		if originURL.Valid {
			book.OriginURL, err = url.Parse(originURL.String)
			if err != nil {
				return fmt.Errorf("parse origin url: %w", err)
			}
		}

		book.Name = name.String
		book.PageCount = int(pageCount.Int32)
		book.DeletedAt = deletedAt.Time
		book.VerifiedAt = verifiedAt.Time

		book.Calc = core.BookCalc{
			CalcPageCount:     NilInt64FromDB(calcPageCount),
			CalcFileCount:     NilInt64FromDB(calcFileCount),
			CalcDeadHashCount: NilInt64FromDB(calcDeadHashCount),
			CalcPageSize:      NilInt64FromDB(calcPageSize),
			CalcFileSize:      NilInt64FromDB(calcFileSize),
			CalcDeadHashSize:  NilInt64FromDB(calcDeadHashSize),
			CalculatedAt:      calculatedAt.Time,
			CalcAvgPageSize:   NilInt64FromDB(calcAvgPageSize),
		}

		return nil
	}
}
