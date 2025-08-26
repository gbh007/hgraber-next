package parsingusecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (uc *UseCase) NewMirror(ctx context.Context, mirror parsing.URLMirror) error {
	mirror.ID = uuid.Must(uuid.NewV7())

	return uc.storage.NewMirror(ctx, mirror) //nolint:wrapcheck // нет смысла оборачивать
}

func (uc *UseCase) UpdateMirror(ctx context.Context, mirror parsing.URLMirror) error {
	return uc.storage.UpdateMirror(ctx, mirror) //nolint:wrapcheck // нет смысла оборачивать
}

func (uc *UseCase) DeleteMirror(ctx context.Context, id uuid.UUID) error {
	return uc.storage.DeleteMirror(ctx, id) //nolint:wrapcheck // нет смысла оборачивать
}

func (uc *UseCase) Mirrors(ctx context.Context) ([]parsing.URLMirror, error) {
	return uc.storage.Mirrors(ctx) //nolint:wrapcheck // нет смысла оборачивать
}

func (uc *UseCase) Mirror(ctx context.Context, id uuid.UUID) (parsing.URLMirror, error) {
	return uc.storage.Mirror(ctx, id) //nolint:wrapcheck // нет смысла оборачивать
}
