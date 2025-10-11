package agent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AgentRepo) Agent(ctx context.Context, id uuid.UUID) (core.Agent, error) {
	table := model.AgentTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnID(): id,
		}).
		Limit(1)

	query, args := builder.MustSql()

	result := core.Agent{}
	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(table.Scanner(&result))

	if errors.Is(err, sql.ErrNoRows) {
		return core.Agent{}, core.ErrAgentNotFound
	}

	if err != nil {
		return core.Agent{}, fmt.Errorf("exec query: %w", err)
	}

	return result, nil
}
