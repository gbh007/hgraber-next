package agent

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AgentRepo) NewAgent(ctx context.Context, agent core.Agent) error {
	table := model.AgentTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnID():            agent.ID,
			table.ColumnName():          agent.Name,
			table.ColumnAddr():          agent.Addr.String(),
			table.ColumnToken():         agent.Token,
			table.ColumnCanParse():      agent.CanParse,
			table.ColumnCanParseMulti(): agent.CanParseMulti,
			table.ColumnCanExport():     agent.CanExport,
			table.ColumnHasFS():         agent.HasFS,
			table.ColumnHasHProxy():     agent.HasHProxy,
			table.ColumnPriority():      agent.Priority,
			table.ColumnCreateAt():      agent.CreateAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
