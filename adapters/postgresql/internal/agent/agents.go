package agent

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AgentRepo) Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error) {
	table := model.AgentTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		OrderBy(table.ColumnPriority() + " DESC")

	if filter.CanParse {
		builder = builder.Where(squirrel.Eq{
			table.ColumnCanParse(): true,
		})
	}

	if filter.CanParseMulti {
		builder = builder.Where(squirrel.Eq{
			table.ColumnCanParseMulti(): true,
		})
	}

	if filter.CanExport {
		builder = builder.Where(squirrel.Eq{
			table.ColumnCanExport(): true,
		})
	}

	if filter.HasFS {
		builder = builder.Where(squirrel.Eq{
			table.ColumnHasFS(): true,
		})
	}

	if filter.HasHProxy {
		builder = builder.Where(squirrel.Eq{
			table.ColumnHasHProxy(): true,
		})
	}

	query, args := builder.MustSql()

	result := make([]core.Agent, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		agent := core.Agent{}

		err := rows.Scan(table.Scanner(&agent))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, agent)
	}

	return result, nil
}
