package files

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID) error {
	filepath := s.filepath(fileID)

	err := os.Remove(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return entities.FileNotFoundError
	}

	if err != nil {
		return fmt.Errorf("os remove: %w", err)
	}

	return nil
}
