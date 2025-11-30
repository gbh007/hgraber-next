package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) InsertAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	attrRemapTable := model.AttributeRemapTable

	builder := squirrel.Insert(attrRemapTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			attrRemapTable.ColumnAttr():      ar.Code,
			attrRemapTable.ColumnValue():     ar.Value,
			attrRemapTable.ColumnToAttr():    model.StringToDB(ar.ToCode),
			attrRemapTable.ColumnToValue():   model.StringToDB(ar.ToValue),
			attrRemapTable.ColumnCreatedAt(): ar.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
