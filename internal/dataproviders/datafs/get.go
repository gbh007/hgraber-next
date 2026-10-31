package datafs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/google/uuid"
)

func (s *Storage) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("get", nil, time.Since(startAt))
	}()

	filepath := s.filepath(fileID)

	f, err := os.Open(filepath)

	if os.IsNotExist(err) {
		return nil, core.ErrFileNotFound
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
