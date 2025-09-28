package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func BookColumns() []string {
	return []string{
		"id",
		"name",
		"origin_url",
		"page_count",
		"attributes_parsed",
		"create_at",
		"deleted",
		"deleted_at",
		"verified",
		"verified_at",
		"is_rebuild",

		"calc_page_count",
		"calc_file_count",
		"calc_dead_hash_count",
		"calc_page_size",
		"calc_file_size",
		"calc_dead_hash_size",
		"calculated_at",
		"calc_avg_page_size",
	}
}

func BookScanner(book *core.Book) RowScanner {
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
