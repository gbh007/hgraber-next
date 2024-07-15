package files

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

func (s *Storage) IDs(ctx context.Context) ([]uuid.UUID, error) {
	entries, err := os.ReadDir(s.fsPath)
	if err != nil {
		return nil, fmt.Errorf("local fs: scan dir: %w", err)
	}

	res := make([]uuid.UUID, 0, len(entries))

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		id, err := uuid.Parse(e.Name())
		if err != nil {
			s.logger.WarnContext(
				ctx, "invalid file in file dir",
				slog.String("filename", e.Name()),
			)

			continue
		}

		res = append(res, id)
	}

	return res, nil
}
