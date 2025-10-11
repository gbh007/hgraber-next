package agent

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AgentRepo) UpdateAgent(ctx context.Context, agent core.Agent) error {
	table := model.AgentTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnName():          agent.Name,
			table.ColumnAddr():          agent.Addr.String(),
			table.ColumnToken():         agent.Token,
			table.ColumnCanParse():      agent.CanParse,
			table.ColumnCanParseMulti(): agent.CanParseMulti,
			table.ColumnCanExport():     agent.CanExport,
			table.ColumnHasFS():         agent.HasFS,
			table.ColumnHasHProxy():     agent.HasHProxy,
			table.ColumnPriority():      agent.Priority,
		}).
		Where(squirrel.Eq{
			table.ColumnID(): agent.ID,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrAgentNotFound
	}

	return nil
}
