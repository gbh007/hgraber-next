package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	logger *slog.Logger
	debug  bool

	pool *pgxpool.Pool
	db   *sqlx.DB
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	debug bool,
	dataSourceName string,
	maxConn int32,
) (*Database, error) {
	pgxConfig, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if maxConn > 0 {
		pgxConfig.MaxConns = maxConn
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	db := sqlx.NewDb(stdlib.OpenDBFromPool(dbpool), "pgx")

	err = migrate(ctx, logger, db.DB)
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return &Database{
		logger: logger,
		debug:  debug,
		pool:   dbpool,
		db:     db,
	}, nil
}

func isApplyWithErr(r sql.Result) (bool, error) {
	c, err := r.RowsAffected()
	if err != nil {
		return false, nil
	}

	return c != 0, nil
}

func (d *Database) isApply(ctx context.Context, r sql.Result) bool {
	apply, err := isApplyWithErr(r)

	if err != nil {
		d.logger.ErrorContext(ctx, err.Error())
	}

	return apply
}

func (d *Database) squirrelDebugLog(ctx context.Context, query string, args []any) {
	if !d.debug {
		return
	}

	d.logger.DebugContext(
		ctx, "squirrel build request",
		slog.String("query", query),
		slog.Any("args", args),
	)
}
