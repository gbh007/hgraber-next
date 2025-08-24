package hproxyusecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

var errCanNotParse = errors.New("can't parse")

type storage interface {
	Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error)

	Mirrors(ctx context.Context) ([]parsing.URLMirror, error)
	GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error)

	Attributes(ctx context.Context) ([]core.Attribute, error)
	AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error)
}

type agentSystem interface {
	BooksCheck(ctx context.Context, agentID uuid.UUID, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	PagesCheck(
		ctx context.Context,
		agentID uuid.UUID,
		urls []agentmodel.AgentPageURL,
	) ([]agentmodel.AgentPageCheckResult, error)

	PageLoad(ctx context.Context, agentID uuid.UUID, url agentmodel.AgentPageURL) (io.Reader, error)

	HProxyList(ctx context.Context, agentID uuid.UUID, u url.URL) (hproxymodel.List, error)
	HProxyBook(ctx context.Context, agentID uuid.UUID, u url.URL, pageLimit *int) (hproxymodel.Book, error)
}

type UseCase struct {
	logger *slog.Logger

	storage     storage
	agentSystem agentSystem

	parseBookTimeout time.Duration

	autoRemap    bool
	remapToLower bool
}

func New(
	logger *slog.Logger,
	storage storage,
	agentSystem agentSystem,
	parseBookTimeout time.Duration,
	autoRemap bool,
	remapToLower bool,
) *UseCase {
	return &UseCase{
		logger:           logger,
		storage:          storage,
		agentSystem:      agentSystem,
		parseBookTimeout: parseBookTimeout,
		autoRemap:        autoRemap,
		remapToLower:     remapToLower,
	}
}

func (uc *UseCase) existsInStorage(ctx context.Context, urls []url.URL) ([]uuid.UUID, error) {
	for _, u := range urls {
		// FIXME: нужно сделать более оптимальный метод
		ids, err := uc.storage.GetBookIDsByURL(ctx, u)
		if err != nil {
			return nil, fmt.Errorf("check exists by (%s): %w", u.String(), err)
		}

		if len(ids) > 0 {
			return ids, nil
		}
	}

	return []uuid.UUID{}, nil
}
