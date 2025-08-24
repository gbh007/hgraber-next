package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/agent"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/attribute"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/book"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/file"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/label"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/massload"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"
)

type Database struct {
	*repository.Repository
	*massload.MassloadRepo
	*file.FileRepo
	*agent.AgentRepo
	*attribute.AttributeRepo
	*label.LabelRepo
	*book.BookRepo
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
	repo, err := repository.New(
		ctx,
		logger,
		tracer,
		debugPgx,
		debugSquirrel,
		dataSourceName,
		maxConn,
	)
	if err != nil {
		return nil, fmt.Errorf("init base repo: %w", err)
	}

	return &Database{
		Repository:    repo,
		MassloadRepo:  massload.New(repo),
		FileRepo:      file.New(repo),
		AgentRepo:     agent.New(repo),
		AttributeRepo: attribute.New(repo),
		LabelRepo:     label.New(repo),
		BookRepo:      book.New(repo),
	}, nil
}
