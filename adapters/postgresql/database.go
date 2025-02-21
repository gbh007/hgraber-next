package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"go.opentelemetry.io/otel/trace"
)

type Database struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	pool *pgxpool.Pool
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	tracer trace.Tracer,
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

	// TODO: начать использовать
	// pgxConfig.ConnConfig.Tracer = pgxTracer{
	// 	logger: logger,
	// 	tracer: tracer,
	// 	debug:  debug,
	// }

	dbpool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	err = migrate(ctx, logger, stdlib.OpenDBFromPool(dbpool))
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return &Database{
		logger: logger,
		tracer: tracer,
		debug:  debug,
		pool:   dbpool,
	}, nil
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

/*  TODO: начать использовать
var _ pgx.QueryTracer = (*pgxTracer)(nil)
var _ pgx.BatchTracer = (*pgxTracer)(nil)

type pgxTracer struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool
}

func (t pgxTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	return ctx
}

func (t pgxTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {}

func (t pgxTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	return ctx
}

func (t pgxTracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
}

func (t pgxTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {}
*/
