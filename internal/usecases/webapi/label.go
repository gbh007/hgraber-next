package webapi

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) SetLabel(ctx context.Context, label entities.BookLabel) error {
	label.CreateAt = time.Now().UTC()

	return uc.storage.SetLabel(ctx, label)
}

func (uc *UseCase) DeleteLabel(ctx context.Context, label entities.BookLabel) error {
	return uc.storage.DeleteLabel(ctx, label)
}

func (uc *UseCase) Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error) {
	return uc.storage.Labels(ctx, bookID)
}

func (uc *UseCase) CreateLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error {
	preset.CreatedAt = time.Now().UTC()

	return uc.storage.InsertLabelPreset(ctx, preset)
}

func (uc *UseCase) UpdateLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error {
	preset.UpdatedAt = time.Now().UTC()

	return uc.storage.UpdateLabelPreset(ctx, preset)
}

func (uc *UseCase) DeleteLabelPreset(ctx context.Context, name string) error {
	return uc.storage.DeleteLabelPreset(ctx, name)
}

func (uc *UseCase) LabelPresets(ctx context.Context) ([]entities.BookLabelPreset, error) {
	return uc.storage.LabelPresets(ctx)
}

func (uc *UseCase) LabelPreset(ctx context.Context, name string) (entities.BookLabelPreset, error) {
	return uc.storage.LabelPreset(ctx, name)
}
