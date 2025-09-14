package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"go.opentelemetry.io/otel/trace"
)

const dbName = "postgres"

type MetricProvider interface {
	IncDBActiveRequest(db string)
	DecDBActiveRequest(db string)
	SetDBOpenConnection(db string, n int32)
	RegisterDBRequestDuration(db, stmt string, d time.Duration)
}

type Repository struct {
	Logger *slog.Logger
	Tracer trace.Tracer

	Pool *pgxpool.Pool
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	tracer trace.Tracer,
	metricProvider MetricProvider,
	debugPgx bool,
	dataSourceName string,
	maxConn int32,
) (*Repository, error) {
	pgxConfig, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if maxConn > 0 {
		pgxConfig.MaxConns = maxConn
	}

	pgxConfig.ConnConfig.Tracer = pgxTracer{
		logger:         logger,
		tracer:         tracer,
		metricProvider: metricProvider,
		debug:          debugPgx,
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	go func() {
		for range time.NewTicker(time.Second * 5).C { //nolint:mnd // будет исправлено позднее
			metricProvider.SetDBOpenConnection(dbName, dbpool.Stat().TotalConns())
		}
	}()

	err = migrate(ctx, logger, stdlib.OpenDBFromPool(dbpool))
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return &Repository{
		Logger: logger,
		Tracer: tracer,
		Pool:   dbpool,
	}, nil
}
