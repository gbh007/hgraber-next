package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
)

var (
	_ pgx.QueryTracer = (*pgxTracer)(nil)
	_ pgx.BatchTracer = (*pgxTracer)(nil)
)

type (
	requestCtxKey      struct{}
	batchRequestCtxKey struct{}
)

type requestInfo struct {
	stmt    string
	startAt time.Time
}

type batchRequestInfo struct {
	startAt time.Time
}

type pgxTracer struct {
	logger         *slog.Logger
	metricProvider MetricProvider
	debug          bool
}

func (t pgxTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	t.metricProvider.IncDBActiveRequest(dbName)

	ctx = context.WithValue(ctx, requestCtxKey{}, requestInfo{
		stmt:    data.SQL,
		startAt: time.Now(),
	})

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
	v, ok := ctx.Value(requestCtxKey{}).(requestInfo)
	if stmt, keep := filterStmt(v.stmt); ok && keep {
		t.metricProvider.RegisterDBRequestDuration(dbName, stmt, time.Since(v.startAt))
	}

	t.metricProvider.DecDBActiveRequest(dbName)
}

func (t pgxTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	queries := lo.Map(data.Batch.QueuedQueries, func(q *pgx.QueuedQuery, _ int) string {
		if q == nil {
			return ""
		}

		return q.SQL
	})

	for range data.Batch.QueuedQueries {
		t.metricProvider.IncDBActiveRequest(dbName)
	}

	ctx = context.WithValue(ctx, batchRequestCtxKey{}, batchRequestInfo{
		startAt: time.Now(),
	})

	if t.debug {
		t.logger.DebugContext(
			ctx, "pgx batch",
			slog.Any("queries", queries),
		)
	}

	return ctx
}

func (t pgxTracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
	if t.debug {
		t.logger.DebugContext(
			ctx, "pgx batch query",
			slog.String("query", data.SQL),
			slog.Any("args", data.Args),
		)
	}

	v, ok := ctx.Value(batchRequestCtxKey{}).(batchRequestInfo)
	if stmt, keep := filterStmt(data.SQL); ok && keep {
		t.metricProvider.RegisterDBRequestDuration(dbName, stmt, time.Since(v.startAt))
	}

	t.metricProvider.DecDBActiveRequest(dbName)
}

func (t pgxTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
}
