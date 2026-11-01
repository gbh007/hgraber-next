package repository

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log/slog"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

var _ goose.Logger = (*slogGooseAdapter)(nil)

type slogGooseAdapter struct {
	logger *slog.Logger
}

func (a slogGooseAdapter) Fatalf(format string, v ...any) {
	a.logger.Error(fmt.Sprintf(format, v...)) //nolint:sloglint // особенность библиотеки goose
}

func (a slogGooseAdapter) Printf(format string, v ...any) {
	a.logger.Info(fmt.Sprintf(format, v...)) //nolint:sloglint // особенность библиотеки goose
}

func migrate(ctx context.Context, logger *slog.Logger, db *sql.DB) error {
	goose.SetBaseFS(migrationsFS)

	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	goose.SetLogger(slogGooseAdapter{
		logger: logger,
	})

	err = goose.UpContext(
		ctx, db, "migrations",
		goose.WithNoColor(true),
	)
	if err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	return nil
}
