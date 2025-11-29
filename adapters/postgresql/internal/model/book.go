package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var BookTable = Book{baseTable: baseTable{name: "books"}}

type Book struct {
	baseTable
}

func (b Book) WithPrefix(pf string) Book {
	return Book{
		baseTable: b.withPrefix(pf),
	}
}

func (b Book) ColumnID() string                { return b.prefix + "id" }
func (b Book) ColumnName() string              { return b.prefix + "name" }
func (b Book) ColumnOriginURL() string         { return b.prefix + "origin_url" }
func (b Book) ColumnPageCount() string         { return b.prefix + "page_count" }
func (b Book) ColumnAttributesParsed() string  { return b.prefix + "attributes_parsed" }
func (b Book) ColumnCreateAt() string          { return b.prefix + "create_at" }
func (b Book) ColumnDeleted() string           { return b.prefix + "deleted" }
func (b Book) ColumnDeletedAt() string         { return b.prefix + "deleted_at" }
func (b Book) ColumnVerified() string          { return b.prefix + "verified" }
func (b Book) ColumnVerifiedAt() string        { return b.prefix + "verified_at" }
func (b Book) ColumnIsRebuild() string         { return b.prefix + "is_rebuild" }
func (b Book) ColumnCalcPageCount() string     { return b.prefix + "calc_page_count" }
func (b Book) ColumnCalcFileCount() string     { return b.prefix + "calc_file_count" }
func (b Book) ColumnCalcDeadHashCount() string { return b.prefix + "calc_dead_hash_count" }
func (b Book) ColumnCalcPageSize() string      { return b.prefix + "calc_page_size" }
func (b Book) ColumnCalcFileSize() string      { return b.prefix + "calc_file_size" }
func (b Book) ColumnCalcDeadHashSize() string  { return b.prefix + "calc_dead_hash_size" }
func (b Book) ColumnCalculatedAt() string      { return b.prefix + "calculated_at" }
func (b Book) ColumnCalcAvgPageSize() string   { return b.prefix + "calc_avg_page_size" }

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
