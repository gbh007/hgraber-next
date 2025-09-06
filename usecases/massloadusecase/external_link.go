package massloadusecase

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (uc *UseCase) CreateMassloadExternalLink(
	ctx context.Context,
	massloadID int,
	link massloadmodel.ExternalLink,
) error {
	link.CreatedAt = time.Now()

	err := uc.storage.CreateMassloadExternalLink(ctx, massloadID, link)
	if err != nil {
		return fmt.Errorf("storage create: %w", err)
	}

	return nil
}

func (uc *UseCase) UpdateMassloadExternalLink(
	ctx context.Context,
	massloadID int,
	link massloadmodel.ExternalLink,
) error {
	link.UpdatedAt = time.Now()

	err := uc.storage.UpdateMassloadExternalLink(ctx, massloadID, link)
	if err != nil {
		return fmt.Errorf("storage update: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteMassloadExternalLink(ctx context.Context, massloadID int, u url.URL) error {
	err := uc.storage.DeleteMassloadExternalLink(ctx, massloadID, u)
	if err != nil {
		return fmt.Errorf("storage delete: %w", err)
	}

	return nil
}
