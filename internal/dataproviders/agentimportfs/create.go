package agentimportfs

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
)

func (s *Storage) Create(ctx context.Context, data agentdto.ImportData) error {
	_, err := s.create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateImport(ctx context.Context, data agentdto.ImportData) (string, error) {
	return s.create(ctx, data)
}

func (s *Storage) create(ctx context.Context, data agentdto.ImportData) (string, error) {
	startAt := time.Now()
	defer func() {
		s.metricProvider.RegisterFSActionTime("create_import", nil, time.Since(startAt))
	}()

	var (
		relativePath, absolutePath string
		err                        error
	)

	if s.limitOnFolder > 0 {
		s.fsLimitMutex.Lock()
		// TODO: подумать над многопоточным способом (включая необходимость)
		defer s.fsLimitMutex.Unlock()

		relativePath, absolutePath, err = s.filepathWithLimits(data.BookID, data.BookName)
		if err != nil {
			return "", fmt.Errorf("get filepath with limits: %w", err)
		}
	} else {
		relativePath, absolutePath = s.filepath(data.BookID, data.BookName)
	}

	f, err := os.Create(absolutePath)
	if err != nil {
		return "", fmt.Errorf("create file: %w", err)
	}

	_, err = io.Copy(f, data.Body)
	if err != nil {
		fileCloseErr := f.Close()
		if fileCloseErr != nil {
			s.logger.ErrorContext(ctx, "close on write error", slog.Any("error", fileCloseErr))
		}

		return "", fmt.Errorf("write file: %w", err)
	}

	err = f.Close()
	if err != nil {
		return "", fmt.Errorf("close file: %w", err)
	}

	return relativePath, nil
}
