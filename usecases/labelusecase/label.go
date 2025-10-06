package labelusecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) SetLabel(ctx context.Context, label core.BookLabel) error {
	label.CreateAt = time.Now().UTC()

	return uc.storage.SetLabel(ctx, label) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) DeleteLabel(ctx context.Context, label core.BookLabel) error {
	return uc.storage.DeleteLabel(ctx, label) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error) {
	return uc.storage.Labels(ctx, bookID) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) CreateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error {
	preset.CreatedAt = time.Now().UTC()

	return uc.storage.InsertLabelPreset(ctx, preset) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) UpdateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error {
	preset.UpdatedAt = time.Now().UTC()

	return uc.storage.UpdateLabelPreset(ctx, preset) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) DeleteLabelPreset(ctx context.Context, name string) error {
	return uc.storage.DeleteLabelPreset(ctx, name) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) LabelPresets(ctx context.Context) ([]core.BookLabelPreset, error) {
	return uc.storage.LabelPresets(ctx) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) LabelPreset(ctx context.Context, name string) (core.BookLabelPreset, error) {
	return uc.storage.LabelPreset(ctx, name) //nolint:wrapcheck // обвязка не требуется
}
