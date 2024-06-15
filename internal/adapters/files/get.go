package files

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

func (s *Storage) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	filepath := s.filepath(fileID)

	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read all: %w", err)
	}

	return bytes.NewReader(data), nil
}
