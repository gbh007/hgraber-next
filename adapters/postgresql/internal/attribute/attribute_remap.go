package attribute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributeRemap(ctx context.Context, code, value string) (core.AttributeRemap, error) {
	builder := squirrel.Select(model.AttributeRemapColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("attribute_remaps").
		Where(squirrel.Eq{
			"attr":  code,
			"value": value,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return core.AttributeRemap{}, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	row := repo.Pool.QueryRow(ctx, query, args...)

	ar := core.AttributeRemap{}

	err = row.Scan(model.AttributeRemapScanner(&ar))
	if errors.Is(err, sql.ErrNoRows) {
		return core.AttributeRemap{}, core.AttributeRemapNotFoundError
	}

	if err != nil {
		return core.AttributeRemap{}, fmt.Errorf("scan row: %w", err)
	}

	return ar, nil
}
