package massloadusecase

import (
	"context"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

type storage interface {
	CreateMassload(ctx context.Context, ml massloadmodel.Massload) (int, error)
	UpdateMassload(ctx context.Context, ml massloadmodel.Massload) error
	UpdateMassloadSize(ctx context.Context, ml massloadmodel.Massload) error
	Massload(ctx context.Context, id int) (massloadmodel.Massload, error)
	Massloads(ctx context.Context, filter massloadmodel.Filter) ([]massloadmodel.Massload, error)
	DeleteMassload(ctx context.Context, id int) error

	MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error)

	CreateMassloadExternalLink(ctx context.Context, id int, link massloadmodel.ExternalLink) error
	MassloadExternalLinks(ctx context.Context, id int) ([]massloadmodel.ExternalLink, error)
	DeleteMassloadExternalLink(ctx context.Context, id int, u url.URL) error

	CreateMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error
	MassloadAttributes(ctx context.Context, id int) ([]massloadmodel.Attribute, error)
	DeleteMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error

	MassloadsAttributes(ctx context.Context) ([]massloadmodel.Attribute, error)
	UpdateMassloadAttributeSize(ctx context.Context, attr massloadmodel.Attribute) error

	AttributesPageSize(ctx context.Context, attrs map[string][]string) (int64, error)
	AttributesFileSize(ctx context.Context, attrs map[string][]string) (int64, error)
}

type UseCase struct {
	logger *slog.Logger

	storage storage
}

func New(
	logger *slog.Logger,
	storage storage,
) *UseCase {
	return &UseCase{
		logger:  logger,
		storage: storage,
	}
}
