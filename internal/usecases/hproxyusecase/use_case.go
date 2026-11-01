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

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/internal/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

var errCanNotParse = errors.New("can't parse")

type storage interface {
	Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error)

	Mirrors(ctx context.Context) ([]parsing.URLMirror, error)
	GetBookIDsByURL(ctx context.Context, u url.URL) ([]uuid.UUID, error)

	Attributes(ctx context.Context) ([]core.Attribute, error)
	AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error)

	MassloadsByAttribute(ctx context.Context, code, value string) ([]massloadmodel.Massload, error)
	MassloadsByExternalLink(ctx context.Context, u url.URL) ([]massloadmodel.Massload, error)
}

type agentSystem interface {
	BooksCheck(ctx context.Context, agentID uuid.UUID, urls []url.URL) ([]serveragentdto.AgentBookCheckResult, error)
	PageLoad(ctx context.Context, agentID uuid.UUID, u serveragentdto.AgentPageURL) (io.Reader, error)

	ParseList(ctx context.Context, agentID uuid.UUID, u url.URL, onePage bool) (hproxymodel.List, error)
	BookParse(ctx context.Context, agentID uuid.UUID, u url.URL, pageLimit *int) (hproxymodel.Book, error)
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
		// TODO: нужно сделать более оптимальный метод
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
