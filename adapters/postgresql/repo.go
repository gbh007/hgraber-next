package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/agent"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/attribute"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/book"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/deadhash"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/file"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/label"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/massload"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/other"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/page"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"
	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/urlmirror"
)

type Repo struct {
	*massload.MassloadRepo
	*file.FileRepo
	*agent.AgentRepo
	*attribute.AttributeRepo
	*label.LabelRepo
	*book.BookRepo
	*deadhash.DeadHashRepo
	*urlmirror.URLMirrorRepo
	*page.PageRepo
	*other.OtherRepo

	// Переопределение для того чтобы реальные не были доступны из вне.
	Repository, Logger, Tracer, Pool struct{} //nolint:revive // это объект пустышка
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	tracer trace.Tracer,
	debugPgx bool,
	debugSquirrel bool,
	dataSourceName string,
	maxConn int32,
) (*Repo, error) {
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

	return &Repo{
		MassloadRepo:  massload.New(repo),
		FileRepo:      file.New(repo),
		AgentRepo:     agent.New(repo),
		AttributeRepo: attribute.New(repo),
		LabelRepo:     label.New(repo),
		BookRepo:      book.New(repo),
		DeadHashRepo:  deadhash.New(repo),
		URLMirrorRepo: urlmirror.New(repo),
		PageRepo:      page.New(repo),
		OtherRepo:     other.New(repo),
	}, nil
}
