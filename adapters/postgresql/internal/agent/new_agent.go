package agent

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AgentRepo) NewAgent(ctx context.Context, agent core.Agent) error {
	builder := squirrel.Insert("agents").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"id":              agent.ID,
			"name":            agent.Name,
			"addr":            agent.Addr.String(),
			"token":           agent.Token,
			"can_parse":       agent.CanParse,
			"can_parse_multi": agent.CanParseMulti,
			"can_export":      agent.CanExport,
			"has_fs":          agent.HasFS,
			"has_hproxy":      agent.HasHProxy,
			"priority":        agent.Priority,
			"create_at":       agent.CreateAt,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
