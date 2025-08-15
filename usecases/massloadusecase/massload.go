package massloadusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (uc *UseCase) CreateMassload(ctx context.Context, ml massloadmodel.Massload) (int, error) {
	ml.CreatedAt = time.Now()

	id, err := uc.storage.CreateMassload(ctx, ml)
	if err != nil {
		return 0, fmt.Errorf("storage create: %w", err)
	}

	return id, nil
}

func (uc *UseCase) UpdateMassload(ctx context.Context, ml massloadmodel.Massload) error {
	ml.UpdatedAt = time.Now()

	err := uc.storage.UpdateMassload(ctx, ml)
	if err != nil {
		return fmt.Errorf("storage update: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteMassload(ctx context.Context, id int) error {
	err := uc.storage.DeleteMassload(ctx, id)
	if err != nil {
		return fmt.Errorf("storage delete: %w", err)
	}

	return nil
}

func (uc *UseCase) Massload(ctx context.Context, id int) (massloadmodel.Massload, error) {
	ml, err := uc.storage.Massload(ctx, id)
	if err != nil {
		return massloadmodel.Massload{}, fmt.Errorf("storage get massload: %w", err)
	}

	ml.Attributes, err = uc.storage.MassloadAttributes(ctx, id)
	if err != nil {
		return massloadmodel.Massload{}, fmt.Errorf("storage get attributes: %w", err)
	}

	ml.ExternalLinks, err = uc.storage.MassloadExternalLinks(ctx, id)
	if err != nil {
		return massloadmodel.Massload{}, fmt.Errorf("storage get external links: %w", err)
	}

	return ml, nil
}

func (uc *UseCase) Massloads(ctx context.Context) ([]massloadmodel.Massload, error) {
	mls, err := uc.storage.Massloads(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage get massloads: %w", err)
	}

	return mls, nil
}
