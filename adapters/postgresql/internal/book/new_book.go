package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) NewBook(ctx context.Context, book core.Book) error {
	builder := squirrel.Insert("books").
		PlaceholderFormat(squirrel.Dollar).SetMap(
		map[string]any{
			"id":                book.ID,
			"name":              model.StringToDB(book.Name),
			"origin_url":        model.URLToDB(book.OriginURL),
			"page_count":        model.Int32ToDB(book.PageCount),
			"attributes_parsed": book.AttributesParsed,
			"verified":          book.Verified,
			"verified_at":       model.TimeToDB(book.VerifiedAt),
			"is_rebuild":        book.IsRebuild,
			"create_at":         book.CreateAt,
		},
	)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("storage: build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("storage: exec query: %w", err)
	}

	return nil
}
