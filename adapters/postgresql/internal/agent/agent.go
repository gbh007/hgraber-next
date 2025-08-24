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
	builder := squirrel.Select(model.AgentColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("agents").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return core.Agent{}, fmt.Errorf("storage: build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	result := core.Agent{}
	row := repo.Pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.AgentScanner(&result))

	if errors.Is(err, sql.ErrNoRows) {
		return core.Agent{}, core.AgentNotFoundError
	}

	if err != nil {
		return core.Agent{}, fmt.Errorf("exec query: %w", err)
	}

	return result, nil
}
