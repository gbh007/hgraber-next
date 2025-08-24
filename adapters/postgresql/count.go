package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

//nolint:lll,gocognit,cyclop,funlen // будет исправлено позднее
func (d *Database) SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error) {
	systemSize := systemmodel.SystemSizeInfo{
		FileCountByFS:         make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		UnhashedFileCountByFS: make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		InvalidFileCountByFS:  make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		DetachedFileCountByFS: make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		PageFileSizeByFS:      make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		FileSizeByFS:          make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
	}
	batch := &pgx.Batch{}

	// Книги

	batch.Queue(`SELECT COUNT(*) FROM books;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.BookCount)
		if err != nil {
			return fmt.Errorf("get book count : %w", err)
		}

		return nil
	})

	batch.Queue(`SELECT COUNT(*) FROM books WHERE deleted = FALSE AND page_count IS NOT NULL AND NOT EXISTS (SELECT 1 FROM pages WHERE book_id = books.id AND pages.downloaded = FALSE);`).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.DownloadedBookCount)
			if err != nil {
				return fmt.Errorf("get downloaded book count : %w", err)
			}

			return nil
		})

	batch.Queue(`SELECT COUNT(*) FROM books WHERE deleted = FALSE AND verified = TRUE AND page_count IS NOT NULL AND NOT EXISTS (SELECT 1 FROM pages WHERE book_id = books.id AND pages.downloaded = FALSE);`).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.VerifiedBookCount)
			if err != nil {
				return fmt.Errorf("get book verified count : %w", err)
			}

			return nil
		})

	batch.Queue(`SELECT COUNT(*) FROM books WHERE is_rebuild = TRUE;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.RebuildedBookCount)
		if err != nil {
			return fmt.Errorf("get book rebuilded count: %w", err)
		}

		return nil
	})

	batch.Queue(`SELECT COUNT(*) FROM books WHERE (name IS NULL OR page_count IS NULL OR attributes_parsed = FALSE) AND origin_url IS NOT NULL AND deleted = FALSE AND is_rebuild = FALSE;`).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.BookUnparsedCount)
			if err != nil {
				return fmt.Errorf("get book unparsed count: %w", err)
			}

			return nil
		})

	batch.Queue(`SELECT COUNT(*) FROM books WHERE deleted = TRUE;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.DeletedBookCount)
		if err != nil {
			return fmt.Errorf("get book deleted count: %w", err)
		}

		return nil
	})

	// Страницы

	batch.Queue(`SELECT COUNT(*) FROM pages;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.PageCount)
		if err != nil {
			return fmt.Errorf("get page count: %w", err)
		}

		return nil
	})

	batch.Queue(`SELECT COUNT(*) FROM pages WHERE downloaded = FALSE;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.PageUnloadedCount)
		if err != nil {
			return fmt.Errorf("get unloaded page count: %w", err)
		}

		return nil
	})

	batch.Queue(`SELECT COUNT(*) FROM pages WHERE file_id IS NULL;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.PageWithoutBodyCount)
		if err != nil {
			return fmt.Errorf("get page without body count: %w", err)
		}

		return nil
	})

	batch.Queue(`SELECT COUNT(*) FROM deleted_pages;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.DeletedPageCount)
		if err != nil {
			return fmt.Errorf("get deleted page count: %w", err)
		}

		return nil
	})

	// Файлы

	batch.Queue(`SELECT COUNT(*) FROM dead_hashes;`).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&systemSize.DeadHashCount)
		if err != nil {
			return fmt.Errorf("get dead hash count: %w", err)
		}

		return nil
	})

	batch.Queue(`SELECT COUNT(*), fs_id FROM files GROUP BY fs_id;`).Query(func(rows pgx.Rows) error {
		defer rows.Close()

		for rows.Next() {
			var (
				count sql.NullInt64
				fsID  uuid.NullUUID
			)

			err := rows.Scan(&count, &fsID)
			if err != nil {
				return fmt.Errorf("get file count: %w", err)
			}

			systemSize.FileCountByFS[fsID.UUID] = count.Int64
		}

		return nil
	})

	batch.Queue(`SELECT COUNT(*), fs_id FROM files WHERE md5_sum IS NULL OR sha256_sum IS NULL OR "size" IS NULL GROUP BY fs_id;`).
		Query(func(rows pgx.Rows) error {
			defer rows.Close()

			for rows.Next() {
				var (
					count sql.NullInt64
					fsID  uuid.NullUUID
				)

				err := rows.Scan(&count, &fsID)
				if err != nil {
					return fmt.Errorf("get unhashed file count: %w", err)
				}

				systemSize.UnhashedFileCountByFS[fsID.UUID] = count.Int64
			}

			return nil
		})

	batch.Queue(`SELECT COUNT(*), fs_id FROM files WHERE invalid_data = TRUE GROUP BY fs_id;`).
		Query(func(rows pgx.Rows) error {
			defer rows.Close()

			for rows.Next() {
				var (
					count sql.NullInt64
					fsID  uuid.NullUUID
				)

				err := rows.Scan(&count, &fsID)
				if err != nil {
					return fmt.Errorf("get invalid file count: %w", err)
				}

				systemSize.InvalidFileCountByFS[fsID.UUID] = count.Int64
			}

			return nil
		})

	batch.Queue(`SELECT COUNT(*), fs_id FROM files WHERE NOT EXISTS (SELECT 1 FROM pages WHERE pages.file_id = files.id) GROUP BY fs_id;`).
		Query(func(rows pgx.Rows) error {
			defer rows.Close()

			for rows.Next() {
				var (
					count sql.NullInt64
					fsID  uuid.NullUUID
				)

				err := rows.Scan(&count, &fsID)
				if err != nil {
					return fmt.Errorf("get detached file count: %w", err)
				}

				systemSize.DetachedFileCountByFS[fsID.UUID] = count.Int64
			}

			return nil
		})

	batch.Queue(`SELECT SUM(f."size"), fs_id FROM pages AS p LEFT JOIN files AS f ON p.file_id = f.id WHERE f."size" IS NOT NULL GROUP BY f.fs_id;`).
		Query(func(rows pgx.Rows) error {
			defer rows.Close()

			for rows.Next() {
				var (
					size sql.NullInt64
					fsID uuid.NullUUID
				)

				err := rows.Scan(&size, &fsID)
				if err != nil {
					return fmt.Errorf("get page file size: %w", err)
				}

				systemSize.PageFileSizeByFS[fsID.UUID] = size.Int64
			}

			return nil
		})

	batch.Queue(`SELECT SUM("size"), fs_id FROM files WHERE "size" IS NOT NULL GROUP BY fs_id;`).
		Query(func(rows pgx.Rows) error {
			defer rows.Close()

			for rows.Next() {
				var (
					size sql.NullInt64
					fsID uuid.NullUUID
				)

				err := rows.Scan(&size, &fsID)
				if err != nil {
					return fmt.Errorf("get file size: %w", err)
				}

				systemSize.FileSizeByFS[fsID.UUID] = size.Int64
			}

			return nil
		})

	batchResult := d.Pool.SendBatch(ctx, batch)

	err := batchResult.Close()
	if err != nil {
		return systemmodel.SystemSizeInfo{}, fmt.Errorf("batch close: %w", err)
	}

	return systemSize, nil
}

//nolint:revive // будет исправлено позднее
