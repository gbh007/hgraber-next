package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/pkg"
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
	tracer         trace.Tracer
	metricProvider MetricProvider
	debug          bool
}

func (t pgxTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	ctx, span := t.tracer.Start(ctx, "pgx query")
	span.SetAttributes(
		attribute.String("pgx.query", data.SQL),
	)

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
	span := trace.SpanFromContext(ctx)
	span.End()

	v, ok := ctx.Value(requestCtxKey{}).(requestInfo)
	if stmt, keep := filterStmt(v.stmt); ok && keep {
		t.metricProvider.RegisterDBRequestDuration(dbName, stmt, time.Since(v.startAt))
	}

	t.metricProvider.DecDBActiveRequest(dbName)
}

func (t pgxTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	queries := pkg.Map(data.Batch.QueuedQueries, func(q *pgx.QueuedQuery) string {
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

	v, ok := ctx.Value(batchRequestCtxKey{}).(batchRequestInfo)
	if stmt, keep := filterStmt(data.SQL); ok && keep {
		t.metricProvider.RegisterDBRequestDuration(dbName, stmt, time.Since(v.startAt))
	}

	t.metricProvider.DecDBActiveRequest(dbName)
}

func (t pgxTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
	span := trace.SpanFromContext(ctx)
	span.End()
}
