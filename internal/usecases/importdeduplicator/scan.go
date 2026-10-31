package importdeduplicator

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

const minEntryPercentage = 0.9999 // Считаем допустимой погрешностью 1 страницу на 10 000.

func (uc *UseCase) ScanZips(ctx context.Context) error {
	relativePaths, err := uc.importFS.AllZips(ctx)
	if err != nil {
		return fmt.Errorf("import fs: scan all zip: %w", err)
	}

	err = uc.storage.TruncateMissing(ctx)
	if err != nil {
		return fmt.Errorf("import fs: truncate missing: %w", err)
	}

	for i, relativePath := range relativePaths {
		uc.logger.DebugContext(
			ctx, "start match archive",
			slog.Int("current", i+1),
			slog.Int("total", len(relativePaths)),
			slog.String("path", relativePath),
		)

		c, err := uc.storage.ImportedCountByRelativePath(ctx, relativePath)
		if err != nil {
			return fmt.Errorf("import fs: get imported count (%s): %w", relativePath, err)
		}

		if c > 0 {
			continue
		}

		body, err := uc.importFS.Get(ctx, relativePath)
		if err != nil {
			return fmt.Errorf("import fs: get zip body (%s): %w", relativePath, err)
		}

		matches, err := uc.masterAPI.DeduplicateArchive(ctx, body)
		if err != nil {
			return fmt.Errorf("master api match (%s): %w", relativePath, err)
		}

		var (
			matched            bool
			maxEntryPercentage float64
		)

		for _, match := range matches {
			if match.EntryPercentage > minEntryPercentage &&
				match.ReverseEntryPercentage > minEntryPercentage {
				err = uc.storage.CreateImport(ctx, entities.ImportInfo{
					BookID:     match.TargetBookID,
					BookURL:    match.OriginBookURL,
					FSPath:     relativePath,
					ImportedAt: time.Now().UTC(),
				})
				if err != nil {
					return fmt.Errorf("storage create import info (%s): %w", relativePath, err)
				}

				matched = true
			}

			if match.EntryPercentage > maxEntryPercentage {
				maxEntryPercentage = match.EntryPercentage
			}
		}

		if !matched {
			err = uc.storage.CreateMissing(ctx, relativePath, maxEntryPercentage)
			if err != nil {
				return fmt.Errorf("storage create missing info (%s): %w", relativePath, err)
			}
		}
	}

	return nil
}
