package massloadusecase

import (
	"context"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type tmpStorage interface {
	SaveTask(task systemmodel.RunnableTask)
}

type storage interface {
	CreateMassload(ctx context.Context, ml massloadmodel.Massload) (int, error)
	UpdateMassload(ctx context.Context, ml massloadmodel.Massload) error
	UpdateMassloadSize(ctx context.Context, ml massloadmodel.Massload) error
	UpdateMassloadCounts(ctx context.Context, ml massloadmodel.Massload) error
	Massload(ctx context.Context, id int) (massloadmodel.Massload, error)
	Massloads(ctx context.Context, filter massloadmodel.Filter) ([]massloadmodel.Massload, error)
	DeleteMassload(ctx context.Context, id int) error

	MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error)
	MassloadFlag(ctx context.Context, code string) (massloadmodel.Flag, error)
	CreateMassloadFlag(ctx context.Context, flag massloadmodel.Flag) error
	UpdateMassloadFlag(ctx context.Context, flag massloadmodel.Flag) error
	DeleteMassloadFlag(ctx context.Context, code string) error

	CreateMassloadExternalLink(ctx context.Context, id int, link massloadmodel.ExternalLink) error
	UpdateMassloadExternalLink(ctx context.Context, id int, link massloadmodel.ExternalLink) error
	UpdateMassloadExternalLinkCounts(ctx context.Context, id int, link massloadmodel.ExternalLink) error
	MassloadExternalLinks(ctx context.Context, id int) ([]massloadmodel.ExternalLink, error)
	DeleteMassloadExternalLink(ctx context.Context, id int, u url.URL) error

	CreateMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error
	MassloadAttributes(ctx context.Context, id int) ([]massloadmodel.Attribute, error)
	DeleteMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error

	MassloadsAttributes(ctx context.Context) ([]massloadmodel.Attribute, error)
	UpdateMassloadAttributeSize(ctx context.Context, attr massloadmodel.Attribute) error

	AttributesPageSize(ctx context.Context, attrs map[string][]string) (core.SizeWithCount, error)
	AttributesFileSize(ctx context.Context, attrs map[string][]string) (core.SizeWithCount, error)

	BookCount(ctx context.Context, filter core.BookFilter) (int, error)
	Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error)
	Mirrors(ctx context.Context) ([]parsing.URLMirror, error)
	GetBookIDsByURL(ctx context.Context, u url.URL) ([]uuid.UUID, error)
}

type agentSystem interface {
	BooksCheck(ctx context.Context, agentID uuid.UUID, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	HProxyList(ctx context.Context, agentID uuid.UUID, u url.URL) (hproxymodel.List, error)
}

type UseCase struct {
	logger *slog.Logger

	storage     storage
	tmpStorage  tmpStorage
	agentSystem agentSystem
}

func New(
	logger *slog.Logger,
	storage storage,
	tmpStorage tmpStorage,
	agentSystem agentSystem,
) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		tmpStorage:  tmpStorage,
		agentSystem: agentSystem,
	}
}
