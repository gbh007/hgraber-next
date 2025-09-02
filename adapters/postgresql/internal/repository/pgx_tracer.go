package repository

import (
	"context"
	"log/slog"
	"regexp"
	"time"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/pkg"
)

var (
	_ pgx.QueryTracer   = (*pgxTracer)(nil)
	_ pgx.BatchTracer   = (*pgxTracer)(nil)
	_ pgx.ConnectTracer = (*pgxTracer)(nil)
)

var (
	stmtSpaceRegexp  = regexp.MustCompile(`\s+`)
	stmtValuesRegexp = regexp.MustCompile(`(\(\s?(?:\$\d+,?\s?)+\),)+`)
	stmtOnRegexp     = regexp.MustCompile(`((?:\$\d+,?\s?)+)`)
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

	t.metricProvider.IncDBActiveRequest()

	ctx = context.WithValue(ctx, requestCtxKey{}, requestInfo{
		stmt:    filterStmt(data.SQL),
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
	if ok {
		t.metricProvider.RegisterDBRequestDuration(v.stmt, time.Since(v.startAt))
	}

	t.metricProvider.DecDBActiveRequest()
}

func (t pgxTracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	queries := pkg.Map(data.Batch.QueuedQueries, func(q *pgx.QueuedQuery) string {
		if q == nil {
			return ""
		}

		return q.SQL
	})

	for range data.Batch.QueuedQueries {
		t.metricProvider.IncDBActiveRequest()
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

	// FIXME: проверить что будет корректно отрабатывать (если будут проблемы использовать TraceBatchEnd)
	v, ok := ctx.Value(batchRequestCtxKey{}).(batchRequestInfo)
	if ok {
		t.metricProvider.RegisterDBRequestDuration(filterStmt(data.SQL), time.Since(v.startAt))
	}

	t.metricProvider.DecDBActiveRequest()
}

func (t pgxTracer) TraceBatchEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchEndData) {
	span := trace.SpanFromContext(ctx)
	span.End()
}

func (t pgxTracer) TraceConnectStart(ctx context.Context, data pgx.TraceConnectStartData) context.Context {
	t.metricProvider.IncDBOpenConnection()

	return ctx
}

func (t pgxTracer) TraceConnectEnd(ctx context.Context, data pgx.TraceConnectEndData) {
	t.metricProvider.DecDBOpenConnection()
}

func filterStmt(s string) string {
	s = stmtSpaceRegexp.ReplaceAllString(s, " ")
	s = stmtValuesRegexp.ReplaceAllString(s, "")
	s = stmtOnRegexp.ReplaceAllString(s, "")

	return s
}
