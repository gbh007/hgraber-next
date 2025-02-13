package parsingusecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (uc *UseCase) NewMirror(ctx context.Context, mirror parsing.URLMirror) error {
	mirror.ID = uuid.Must(uuid.NewV7())

	return uc.storage.NewMirror(ctx, mirror)
}

func (uc *UseCase) UpdateMirror(ctx context.Context, mirror parsing.URLMirror) error {
	return uc.storage.UpdateMirror(ctx, mirror)
}

func (uc *UseCase) DeleteMirror(ctx context.Context, id uuid.UUID) error {
	return uc.storage.DeleteMirror(ctx, id)
}

func (uc *UseCase) Mirrors(ctx context.Context) ([]parsing.URLMirror, error) {
	return uc.storage.Mirrors(ctx)
}

func (uc *UseCase) Mirror(ctx context.Context, id uuid.UUID) (parsing.URLMirror, error) {
	return uc.storage.Mirror(ctx, id)
}
