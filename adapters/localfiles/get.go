package localfiles

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (s *Storage) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	filepath := s.filepath(fileID)

	f, err := os.Open(filepath) //nolint:gosec // не применимо

	if os.IsNotExist(err) {
		return nil, core.FileNotFoundError
	}

	if err != nil {
		return nil, fmt.Errorf("local fs: open: %w", err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			s.logger.ErrorContext(ctx, "close file after get", slog.String("err", err.Error()))
		}
	}()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("local fs: read all: %w", err)
	}

	return bytes.NewReader(data), nil
}
