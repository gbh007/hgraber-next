package massloadusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (uc *UseCase) MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error) {
	mls, err := uc.storage.MassloadFlags(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage get massload flags: %w", err)
	}

	return mls, nil
}

func (uc *UseCase) CreateMassloadFlag(ctx context.Context, flag massloadmodel.Flag) error {
	flag.CreatedAt = time.Now()

	err := uc.storage.CreateMassloadFlag(ctx, flag)
	if err != nil {
		return fmt.Errorf("storage create: %w", err)
	}

	return nil
}

func (uc *UseCase) UpdateMassloadFlag(ctx context.Context, flag massloadmodel.Flag) error {
	err := uc.storage.UpdateMassloadFlag(ctx, flag)
	if err != nil {
		return fmt.Errorf("storage update: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteMassloadFlag(ctx context.Context, code string) error {
	err := uc.storage.DeleteMassloadFlag(ctx, code)
	if err != nil {
		return fmt.Errorf("storage delete: %w", err)
	}

	return nil
}

func (uc *UseCase) MassloadFlag(ctx context.Context, code string) (massloadmodel.Flag, error) {
	flag, err := uc.storage.MassloadFlag(ctx, code)
	if err != nil {
		return massloadmodel.Flag{}, fmt.Errorf("storage get: %w", err)
	}

	return flag, nil
}
