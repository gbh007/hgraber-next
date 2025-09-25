package agent

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AgentRepo) UpdateAgent(ctx context.Context, agent core.Agent) error {
	builder := squirrel.Update("agents").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"name":            agent.Name,
			"addr":            agent.Addr.String(),
			"token":           agent.Token,
			"can_parse":       agent.CanParse,
			"can_parse_multi": agent.CanParseMulti,
			"can_export":      agent.CanExport,
			"has_fs":          agent.HasFS,
			"has_hproxy":      agent.HasHProxy,
			"priority":        agent.Priority,
		}).
		Where(squirrel.Eq{
			"id": agent.ID,
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
