package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (d *Database) Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error) {
	raw := make([]model.Agent, 0)

	builder := squirrel.Select("*").
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

	err = d.db.SelectContext(ctx, &raw, query, args...)
	if err != nil {
		return nil, fmt.Errorf("storage: exec query: %w", err)
	}

	result, err := pkg.MapWithError(raw, func(a model.Agent) (core.Agent, error) {
		return a.ToEntity()
	})
	if err != nil {
		return nil, fmt.Errorf("storage: convert: %w", err)
	}

	return result, nil
}

func (d *Database) Agent(ctx context.Context, id uuid.UUID) (core.Agent, error) {
	raw := model.Agent{}

	builder := squirrel.Select("*").
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

	err = d.db.GetContext(ctx, &raw, query, args...)
	if err != nil {
		return core.Agent{}, fmt.Errorf("storage: exec query: %w", err)
	}

	result, err := raw.ToEntity()
	if err != nil {
		return core.Agent{}, fmt.Errorf("storage: convert: %w", err)
	}

	return result, nil
}

func (d *Database) NewAgent(ctx context.Context, agent core.Agent) error {
	builder := squirrel.Insert("agents").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"id":              agent.ID.String(),
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

	_, err = d.db.ExecContext(ctx, query, args...)
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

	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if !d.isApply(ctx, res) {
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

	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if !d.isApply(ctx, res) {
		return core.AgentNotFoundError
	}

	return nil
}
