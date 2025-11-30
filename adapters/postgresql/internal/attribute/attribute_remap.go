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
	attrRemapTable := model.AttributeRemapTable

	builder := squirrel.Select(attrRemapTable.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(attrRemapTable.Name()).
		Where(squirrel.Eq{
			attrRemapTable.ColumnAttr():  code,
			attrRemapTable.ColumnValue(): value,
		}).
		Limit(1)

	query, args := builder.MustSql()

	row := repo.Pool.QueryRow(ctx, query, args...)

	ar := core.AttributeRemap{}

	err := row.Scan(attrRemapTable.Scanner(&ar))
	if errors.Is(err, sql.ErrNoRows) {
		return core.AttributeRemap{}, core.ErrAttributeRemapNotFound
	}

	if err != nil {
		return core.AttributeRemap{}, fmt.Errorf("scan row: %w", err)
	}

	return ar, nil
}
