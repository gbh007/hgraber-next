package files

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

func (s *Storage) Create(ctx context.Context, bookID uuid.UUID, bookName string, body io.Reader) error {
	filepath := s.filepath(bookID, bookName)

	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	_, err = io.Copy(f, body)
	if err != nil {
		fileCloseErr := f.Close()
		if fileCloseErr != nil {
			s.logger.ErrorContext(ctx, "close on write error", slog.Any("error", fileCloseErr))
		}

		return fmt.Errorf("write file: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("close file: %w", err)
	}

	return nil
}
