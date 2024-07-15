package files

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

func (s *Storage) Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error {
	filepath := s.filepath(fileID)

	info, err := os.Stat(filepath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("local fs: check: %w", err)
	}

	if info != nil {
		return fmt.Errorf("local fs: file exists")
	}

	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("local fs: create: %w", err)
	}

	_, err = io.Copy(f, body)
	if err != nil {
		fileCloseErr := f.Close()
		if fileCloseErr != nil {
			s.logger.ErrorContext(ctx, "close on write error", slog.Any("error", fileCloseErr))
		}

		return fmt.Errorf("local fs: write file: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("local fs: close file: %w", err)
	}

	return nil
}
