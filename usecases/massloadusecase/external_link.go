package massloadusecase

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (uc *UseCase) CreateMassloadExternalLink(ctx context.Context, massloadID int, u url.URL) error {
	err := uc.storage.CreateMassloadExternalLink(ctx, massloadID, massloadmodel.ExternalLink{
		URL:       u,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("storage create: %w", err)
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
