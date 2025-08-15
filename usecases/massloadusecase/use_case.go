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
	Massloads(ctx context.Context) ([]massloadmodel.Massload, error)
	DeleteMassload(ctx context.Context, id int) error

	CreateMassloadExternalLink(ctx context.Context, id int, link massloadmodel.MassloadExternalLink) error
	MassloadExternalLinks(ctx context.Context, id int) ([]massloadmodel.MassloadExternalLink, error)
	DeleteMassloadExternalLink(ctx context.Context, id int, u url.URL) error

	CreateMassloadAttribute(ctx context.Context, id int, attr massloadmodel.MassloadAttribute) error
	MassloadAttributes(ctx context.Context, id int) ([]massloadmodel.MassloadAttribute, error)
	DeleteMassloadAttribute(ctx context.Context, id int, attr massloadmodel.MassloadAttribute) error
	UpdateMassloadAttributeSize(ctx context.Context, attr massloadmodel.MassloadAttribute) error
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
