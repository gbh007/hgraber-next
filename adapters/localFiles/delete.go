package localFiles

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID) error {
	filepath := s.filepath(fileID)

	err := os.Remove(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("local fs: %w", core.FileNotFoundError)
	}

	if err != nil {
		return fmt.Errorf("local fs: os remove: %w", err)
	}

	return nil
}
