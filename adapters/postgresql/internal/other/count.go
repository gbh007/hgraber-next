//revive:disable:file-length-limit
package other

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

//nolint:gocognit,cyclop,funlen // будет исправлено позднее
func (repo *OtherRepo) SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error) {
	systemSize := systemmodel.SystemSizeInfo{
		FileCountByFS:         make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		UnhashedFileCountByFS: make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		InvalidFileCountByFS:  make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		DetachedFileCountByFS: make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		PageFileSizeByFS:      make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
		FileSizeByFS:          make(map[uuid.UUID]int64, fsmodel.ApproximateFSCount),
	}
	batch := &pgx.Batch{}

	bookTable := model.BookTable.WithPrefix("b")
	pageTable := model.PageTable.WithPrefix("p")
	deletedPageTable := model.DeletedPageTable.WithPrefix("dp")
	deadHashTable := model.DeadHashTable.WithPrefix("dh")
	fileTable := model.FileTable.WithPrefix("f")

	apply := func(f func() (string, []any)) *pgx.QueuedQuery {
		q, args := f()

		return batch.Queue(q, args...)
	}
	// Queue(query string, arguments ...any)

	// Книги

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(bookTable.NameAlter()).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.BookCount)
			if err != nil {
				return fmt.Errorf("get book count : %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(bookTable.NameAlter()).
			Where(squirrel.Eq{
				bookTable.ColumnDeleted(): false,
			}).
			Where(squirrel.NotEq{
				bookTable.ColumnPageCount(): nil,
			}).
			Where(
				squirrel.Select("1").
					From(pageTable.NameAlter()).
					Where(pageTable.ColumnBookID() + " = " + bookTable.ColumnID()).
					Where(squirrel.Eq{pageTable.ColumnDownloaded(): false}).
					Prefix(" NOT EXISTS ( ").
					Suffix(" ) "),
			).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.DownloadedBookCount)
			if err != nil {
				return fmt.Errorf("get downloaded book count : %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(bookTable.NameAlter()).
			Where(squirrel.Eq{
				bookTable.ColumnDeleted():  false,
				bookTable.ColumnVerified(): true,
			}).
			Where(squirrel.NotEq{
				bookTable.ColumnPageCount(): nil,
			}).
			Where(
				squirrel.Select("1").
					From(pageTable.NameAlter()).
					Where(pageTable.ColumnBookID() + " = " + bookTable.ColumnID()).
					Where(squirrel.Eq{pageTable.ColumnDownloaded(): false}).
					Prefix(" NOT EXISTS ( ").
					Suffix(" ) "),
			).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.VerifiedBookCount)
			if err != nil {
				return fmt.Errorf("get book verified count : %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(bookTable.NameAlter()).
			Where(squirrel.Eq{
				bookTable.ColumnIsRebuild(): true,
			}).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.RebuildedBookCount)
			if err != nil {
				return fmt.Errorf("get book rebuilded count: %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(bookTable.NameAlter()).
			Where(squirrel.Or{
				squirrel.Eq{bookTable.ColumnName(): nil},
				squirrel.Eq{bookTable.ColumnPageCount(): nil},
				squirrel.Eq{bookTable.ColumnAttributesParsed(): false},
			}).
			Where(squirrel.NotEq{
				bookTable.ColumnOriginURL(): nil,
			}).
			Where(squirrel.Eq{
				bookTable.ColumnDeleted():   false,
				bookTable.ColumnIsRebuild(): false,
			}).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.BookUnparsedCount)
			if err != nil {
				return fmt.Errorf("get book unparsed count: %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(bookTable.NameAlter()).
			Where(squirrel.Eq{
				bookTable.ColumnDeleted(): true,
			}).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.DeletedBookCount)
			if err != nil {
				return fmt.Errorf("get book deleted count: %w", err)
			}

			return nil
		})

	// Страницы

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(pageTable.NameAlter()).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.PageCount)
			if err != nil {
				return fmt.Errorf("get page count: %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(pageTable.NameAlter()).
			Where(squirrel.Eq{
				pageTable.ColumnDownloaded(): false,
			}).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.PageUnloadedCount)
			if err != nil {
				return fmt.Errorf("get unloaded page count: %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(pageTable.NameAlter()).
			Where(squirrel.Eq{
				pageTable.ColumnFileID(): nil,
			}).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.PageWithoutBodyCount)
			if err != nil {
				return fmt.Errorf("get page without body count: %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(deletedPageTable.NameAlter()).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.DeletedPageCount)
			if err != nil {
				return fmt.Errorf("get deleted page count: %w", err)
			}

			return nil
		})

	// Файлы

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)").
			PlaceholderFormat(squirrel.Dollar).
			From(deadHashTable.NameAlter()).
			MustSql()
	}).
		QueryRow(func(row pgx.Row) error {
			err := row.Scan(&systemSize.DeadHashCount)
			if err != nil {
				return fmt.Errorf("get dead hash count: %w", err)
			}

			return nil
		})

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)", fileTable.ColumnFSID()).
			PlaceholderFormat(squirrel.Dollar).
			From(fileTable.NameAlter()).
			GroupBy(fileTable.ColumnFSID()).
			MustSql()
	}).
		Query(func(rows pgx.Rows) error {
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

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)", fileTable.ColumnFSID()).
			PlaceholderFormat(squirrel.Dollar).
			From(fileTable.NameAlter()).
			Where(squirrel.Or{
				squirrel.Eq{fileTable.ColumnMd5Sum(): nil},
				squirrel.Eq{fileTable.ColumnSha256Sum(): nil},
				squirrel.Eq{fileTable.ColumnSize(): nil},
			}).
			GroupBy(fileTable.ColumnFSID()).
			MustSql()
	}).
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

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)", fileTable.ColumnFSID()).
			PlaceholderFormat(squirrel.Dollar).
			From(fileTable.NameAlter()).
			Where(squirrel.Eq{fileTable.ColumnInvalidData(): true}).
			GroupBy(fileTable.ColumnFSID()).
			MustSql()
	}).
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

	apply(func() (string, []any) {
		return squirrel.Select("COUNT(*)", fileTable.ColumnFSID()).
			PlaceholderFormat(squirrel.Dollar).
			From(fileTable.NameAlter()).
			Where(
				squirrel.Select("1").
					From(pageTable.NameAlter()).
					Where(pageTable.ColumnFileID() + " = " + fileTable.ColumnID()).
					Prefix(" NOT EXISTS ( ").
					Suffix(" ) "),
			).
			GroupBy(fileTable.ColumnFSID()).
			MustSql()
	}).
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

	apply(func() (string, []any) {
		return squirrel.Select("SUM("+fileTable.ColumnSize()+")", fileTable.ColumnFSID()).
			PlaceholderFormat(squirrel.Dollar).
			From(pageTable.NameAlter()).
			LeftJoin(model.JoinPageAndFile(pageTable, fileTable)).
			Where(squirrel.NotEq{fileTable.ColumnSize(): nil}).
			GroupBy(fileTable.ColumnFSID()).
			MustSql()
	}).
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

	apply(func() (string, []any) {
		return squirrel.Select("SUM("+fileTable.ColumnSize()+")", fileTable.ColumnFSID()).
			PlaceholderFormat(squirrel.Dollar).
			From(fileTable.NameAlter()).
			Where(squirrel.NotEq{fileTable.ColumnSize(): nil}).
			GroupBy(fileTable.ColumnFSID()).
			MustSql()
	}).
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

	batchResult := repo.Pool.SendBatch(ctx, batch)

	err := batchResult.Close()
	if err != nil {
		return systemmodel.SystemSizeInfo{}, fmt.Errorf("batch close: %w", err)
	}

	return systemSize, nil
}
