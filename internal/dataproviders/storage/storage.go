package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2"
)

const dbName = "sqlite"

type metricProvider interface {
	IncDBActiveRequest(db string)
	DecDBActiveRequest(db string)
	SetDBOpenConnection(db string, n int32)
	RegisterDBRequestDuration(db, stmt string, d time.Duration)
}

type Storage struct {
	db *sqlx.DB

	logger         *slog.Logger
	metricProvider metricProvider
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	metricProvider metricProvider,
	path string,
) (*Storage, error) {
	sql.Register("sqlite3WithHooks", sqlhooks.Wrap(&sqlite.Driver{}, hooks{metricProvider: metricProvider}))

	db, err := sqlx.Open("sqlite3WithHooks", path)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	err = migrate(ctx, logger, db.DB)
	if err != nil {
		return nil, fmt.Errorf("migrate db: %w", err)
	}

	go func() {
		for range time.NewTicker(time.Second * 5).C { //nolint:mnd // будет исправлено позднее
			metricProvider.SetDBOpenConnection(dbName, int32(db.Stats().OpenConnections))
		}
	}()

	s := Storage{
		db:             db,
		metricProvider: metricProvider,
		logger:         logger,
	}

	return &s, nil
}

func URLToDB(u *url.URL) sql.NullString {
	if u == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: u.String(),
		Valid:  true,
	}
}
