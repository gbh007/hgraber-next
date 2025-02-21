package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/pkg"
)

type Database struct {
	logger        *slog.Logger
	tracer        trace.Tracer
	debugPgx      bool
	debugSquirrel bool

	pool *pgxpool.Pool
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	tracer trace.Tracer,
	debugPgx bool,
	debugSquirrel bool,
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

	pgxConfig.ConnConfig.Tracer = pgxTracer{
		logger: logger,
		tracer: tracer,
		debug:  debugPgx,
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	err = migrate(ctx, logger, stdlib.OpenDBFromPool(dbpool))
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return &Database{
		logger:        logger,
		tracer:        tracer,
		debugSquirrel: debugSquirrel,
		debugPgx:      debugPgx,
		pool:          dbpool,
	}, nil
}

func (d *Database) squirrelDebugLog(ctx context.Context, query string, args []any) {
	if !d.debugSquirrel {
		return
	}

	d.logger.DebugContext(
		ctx, "squirrel build request",
		slog.String("query", query),
		slog.Any("args", args),
	)
}

var _ pgx.QueryTracer = (*pgxTracer)(nil)
var _ pgx.BatchTracer = (*pgxTracer)(nil)

type pgxTracer struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool
}

func (t pgxTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	ctx, span := t.tracer.Start(ctx, "pgx query")
	span.SetAttributes(
		attribute.String("pgx.query", data.SQL),
	)

	if t.debug {
		t.logger.DebugContext(
			ctx, "pgx query",
			slog.String("query", data.SQL),
			slog.Any("args", data.Args),
		)
	}

	return ctx
}

func (t pgxTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	span := trace.SpanFromContext(ctx)
	span.End()
}

func (t pgxTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	queries := pkg.Map(data.Batch.QueuedQueries, func(q *pgx.QueuedQuery) string {
		if q == nil {
			return ""
		}

		return q.SQL
	})

	ctx, span := t.tracer.Start(ctx, "pgx query")
	span.SetAttributes(
		attribute.StringSlice("pgx.batch.queries", queries),
	)

	if t.debug {
		t.logger.DebugContext(
			ctx, "pgx batch",
			slog.Any("queries", queries),
		)
	}

	return ctx
}

func (t pgxTracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("pgx.batch.query", trace.WithAttributes(

		attribute.String("pgx.batch.query", data.SQL),
	))

	if t.debug {
		t.logger.DebugContext(
			ctx, "pgx batch query",
			slog.String("query", data.SQL),
			slog.Any("args", data.Args),
		)
	}
}

func (t pgxTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
	span := trace.SpanFromContext(ctx)
	span.End()
}
