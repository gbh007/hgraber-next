package datafs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

func (s *Storage) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("get", nil, time.Since(startAt))
	}()

	filepath := s.filepath(fileID)

	f, err := os.Open(filepath)

	if os.IsNotExist(err) {
		return nil, entities.FileNotFoundError
	}

	if err != nil {
		return nil, fmt.Errorf("local fs: open: %w", err)
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("local fs: read all: %w", err)
	}

	return bytes.NewReader(data), nil
}
