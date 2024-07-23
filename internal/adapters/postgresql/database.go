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

type logger interface {
	Logger(ctx context.Context) *slog.Logger
}

type Database struct {
	pool *pgxpool.Pool
	db   *sqlx.DB

	logger logger
}

func New(ctx context.Context, dataSourceName string, logger logger) (*Database, error) {
	dbpool, err := pgxpool.New(ctx, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	db := sqlx.NewDb(stdlib.OpenDBFromPool(dbpool), "pgx")

	err = migrate(ctx, logger, db.DB)
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return &Database{
		pool:   dbpool,
		db:     db,
		logger: logger,
	}, nil
}

func isApplyWithErr(r sql.Result) (bool, error) {
	c, err := r.RowsAffected()
	if err != nil {
		return false, nil
	}

	return c != 0, nil
}

func (storage *Database) isApply(ctx context.Context, r sql.Result) bool {
	apply, err := isApplyWithErr(r)

	if err != nil {
		storage.logger.Logger(ctx).ErrorContext(ctx, err.Error())
	}

	return apply
}
