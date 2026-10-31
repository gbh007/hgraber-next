package storage

import (
	"context"
	"regexp"
	"time"
)

var (
	stmtSpaceRegexp  = regexp.MustCompile(`\s+`)
	stmtValuesRegexp = regexp.MustCompile(`(\(\s?(?:\?,?\s?)+\),)+`)
	stmtOnRegexp     = regexp.MustCompile(`((?:\?,?\s?)+)`)
)

type (
	requestCtxKey struct{}
)

type hooks struct {
	metricProvider metricProvider
}

type requestInfo struct {
	stmt    string
	startAt time.Time
}

func (h hooks) Before(ctx context.Context, query string, args ...any) (context.Context, error) {
	ctx = context.WithValue(ctx, requestCtxKey{}, requestInfo{
		stmt:    filterStmt(query),
		startAt: time.Now(),
	})

	h.metricProvider.IncDBActiveRequest(dbName)

	return ctx, nil
}

func (h hooks) After(ctx context.Context, query string, args ...any) (context.Context, error) {
	v, ok := ctx.Value(requestCtxKey{}).(requestInfo)
	if ok {
		h.metricProvider.RegisterDBRequestDuration(dbName, v.stmt, time.Since(v.startAt))
	}

	h.metricProvider.DecDBActiveRequest(dbName)

	return ctx, nil
}

func filterStmt(s string) string {
	s = stmtSpaceRegexp.ReplaceAllString(s, " ")
	s = stmtValuesRegexp.ReplaceAllString(s, "")
	s = stmtOnRegexp.ReplaceAllString(s, "")

	return s
}
