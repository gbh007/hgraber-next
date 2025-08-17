package massloadusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (uc *UseCase) MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error) {
	mls, err := uc.storage.MassloadFlags(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage get massload flags: %w", err)
	}

	return mls, nil
}
