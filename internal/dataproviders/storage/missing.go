package storage

import (
	"context"
	"time"
)

func (s *Storage) CreateMissing(ctx context.Context, path string, maxEntryPercentage float64) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT INTO missing_infos (relative_path, scanned_at, max_entry_percentage) VALUES (?,?,?) ON CONFLICT DO NOTHING;`,
		path,
		time.Now().UTC().Unix(),
		maxEntryPercentage,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) TruncateMissing(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM missing_infos;`)
	if err != nil {
		return err
	}

	return nil
}
