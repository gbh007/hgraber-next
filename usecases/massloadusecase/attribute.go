package massloadusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (uc *UseCase) CreateMassloadAttribute(ctx context.Context, massloadID int, code, value string) error {
	err := uc.storage.CreateMassloadAttribute(ctx, massloadID, massloadmodel.MassloadAttribute{
		AttrCode:  code,
		AttrValue: value,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("storage create: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteMassloadAttribute(ctx context.Context, massloadID int, code, value string) error {
	err := uc.storage.DeleteMassloadAttribute(ctx, massloadID, massloadmodel.MassloadAttribute{
		AttrCode:  code,
		AttrValue: value,
	})
	if err != nil {
		return fmt.Errorf("storage delete: %w", err)
	}

	return nil
}
