package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func MassloadColumns() []string {
	return []string{
		"id",
		"name",
		"description",
		"flags",
		"page_size",
		"file_size",
		"created_at",
		"updated_at",
	}
}

func MassloadScanner(ml *massloadmodel.Massload) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			description sql.NullString
			pageSize    sql.NullInt64
			fileSize    sql.NullInt64
			updatedAt   sql.NullTime
			// flags       pgtype.Array[string]
		)

		err := rows.Scan(
			&ml.ID,
			&ml.Name,
			&description,
			// &flags,
			&ml.Flags,
			&pageSize,
			&fileSize,
			&ml.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		ml.UpdatedAt = updatedAt.Time
		ml.Description = description.String
		// ml.Flags = flags.Elements

		if pageSize.Valid {
			ml.PageSize = &pageSize.Int64 // TODO: оптимизировать чтобы не уходили в кучу лишние данные.
		}

		if fileSize.Valid {
			ml.FileSize = &fileSize.Int64 // TODO: оптимизировать чтобы не уходили в кучу лишние данные.
		}

		return nil
	}
}

func MassloadExternalLinkColumns() []string {
	return []string{
		"url",
		"created_at",
	}
}

func MassloadExternalLinkScanner(link *massloadmodel.MassloadExternalLink) RowScanner {
	return func(rows pgx.Rows) error {
		var rawUrl string

		err := rows.Scan(
			&rawUrl,
			&link.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		u, err := url.Parse(rawUrl)
		if err != nil {
			return fmt.Errorf("parse url: %w", err)
		}

		link.URL = *u

		return nil
	}
}

func MassloadAttributeColumns() []string {
	return []string{
		"attr_code",
		"attr_value",
		"page_size",
		"file_size",
		"created_at",
		"updated_at",
	}
}

func MassloadAttributeScanner(attr *massloadmodel.MassloadAttribute) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			pageSize  sql.NullInt64
			fileSize  sql.NullInt64
			updatedAt sql.NullTime
		)

		err := rows.Scan(
			&attr.AttrCode,
			&attr.AttrValue,
			&pageSize,
			&fileSize,
			&attr.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		attr.UpdatedAt = updatedAt.Time

		if pageSize.Valid {
			attr.PageSize = &pageSize.Int64 // TODO: оптимизировать чтобы не уходили в кучу лишние данные.
		}

		if fileSize.Valid {
			attr.FileSize = &fileSize.Int64 // TODO: оптимизировать чтобы не уходили в кучу лишние данные.
		}

		return nil
	}
}

func MassloadFlagColumns() []string {
	return []string{
		"code",
		"name",
		"description",
		"created_at",
	}
}

func MassloadFlagScanner(ml *massloadmodel.Flag) RowScanner {
	return func(rows pgx.Rows) error {
		var description sql.NullString

		err := rows.Scan(
			&ml.Code,
			&ml.Name,
			&description,
			&ml.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		ml.Description = description.String

		return nil
	}
}
