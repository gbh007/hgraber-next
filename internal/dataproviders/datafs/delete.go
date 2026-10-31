package datafs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/google/uuid"
)

func (s *Storage) Delete(ctx context.Context, fileID uuid.UUID) error {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("delete", nil, time.Since(startAt))
	}()

	filepath := s.filepath(fileID)

	err := os.Remove(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("local fs: %w", entities.FileNotFoundError)
	}

	if err != nil {
		return fmt.Errorf("local fs: os remove: %w", err)
	}

	return nil
}
