package postgresql

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

func (d *Database) Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error) {
	builder := squirrel.Select(model.AgentColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("agents").OrderBy("priority DESC")

	if filter.CanParse {
		builder = builder.Where(squirrel.Eq{
			"can_parse": true,
		})
	}

	if filter.CanParseMulti {
		builder = builder.Where(squirrel.Eq{
			"can_parse_multi": true,
		})
	}

	if filter.CanExport {
		builder = builder.Where(squirrel.Eq{
			"can_export": true,
		})
	}

	if filter.HasFS {
		builder = builder.Where(squirrel.Eq{
			"has_fs": true,
		})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("storage: build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]core.Agent, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.Agent{}

		err := rows.Scan(model.AgentScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, page)
	}

	return result, nil
}

func (d *Database) Agent(ctx context.Context, id uuid.UUID) (core.Agent, error) {
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

	d.squirrelDebugLog(ctx, query, args)

	result := core.Agent{}
	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.AgentScanner(&result))

	if errors.Is(err, sql.ErrNoRows) {
		return core.Agent{}, core.AgentNotFoundError
	}

	if err != nil {
		return core.Agent{}, fmt.Errorf("exec query :%w", err)
	}

	return result, nil
}

func (d *Database) NewAgent(ctx context.Context, agent core.Agent) error {
	builder := squirrel.Insert("agents").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"id":              agent.ID,
			"name":            agent.Name,
			"addr":            agent.Addr.String(),
			"token":           agent.Token,
			"can_parse":       agent.CanParse,
			"can_parse_multi": agent.CanParseMulti,
			"can_export":      agent.CanExport,
			"has_fs":          agent.HasFS,
			"priority":        agent.Priority,
			"create_at":       agent.CreateAt,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) UpdateAgent(ctx context.Context, agent core.Agent) error {
	builder := squirrel.Update("agents").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"name":            agent.Name,
			"addr":            agent.Addr.String(),
			"token":           agent.Token,
			"can_parse":       agent.CanParse,
			"can_parse_multi": agent.CanParseMulti,
			"can_export":      agent.CanExport,
			"has_fs":          agent.HasFS,
			"priority":        agent.Priority,
		}).
		Where(squirrel.Eq{
			"id": agent.ID.String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.AgentNotFoundError
	}

	return nil
}

func (d *Database) DeleteAgent(ctx context.Context, id uuid.UUID) error {
	builder := squirrel.Delete("agents").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id.String(),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	res, err := d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.AgentNotFoundError
	}

	return nil
}
