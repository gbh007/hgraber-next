package serverdb

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/agent"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/attribute"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/book"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/deadhash"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/file"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/label"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/massload"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/other"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/page"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/repository"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/urlmirror"
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
	Repository, Logger, Tracer, Pool struct{}
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	metricProvider repository.MetricProvider,
	debugPgx bool,
	dataSourceName string,
	maxConn int32,
) (*Repo, error) {
	repo, err := repository.New(
		ctx,
		logger,
		metricProvider,
		debugPgx,
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
