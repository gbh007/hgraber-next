package agentimportapiuc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
)

func (uc *UseCase) Create(ctx context.Context, data agentdto.ImportData) error {
	c, err := uc.storage.ImportedCountByID(ctx, data.BookID)
	if err != nil {
		return fmt.Errorf("check import count by id: %w", err)
	}

	if c > 0 {
		uc.logger.DebugContext(
			ctx, "import already exists",
			slog.String("book_id", data.BookID.String()),
		)

		return nil
	}

	if data.BookURL != nil {
		c, err := uc.storage.ImportedCountByURL(ctx, *data.BookURL)
		if err != nil {
			return fmt.Errorf("check import count by url: %w", err)
		}

		if c > 0 {
			uc.logger.DebugContext(
				ctx, "import already exists",
				slog.String("book_id", data.BookID.String()),
				slog.String("book_url", data.BookURL.String()),
			)

			return nil
		}
	}

	relativePath, err := uc.importFS.CreateImport(ctx, data)
	if err != nil {
		return fmt.Errorf("create file in import fs: %w", err)
	}

	err = uc.storage.CreateImport(ctx, agentdto.ImportInfo{
		BookID:     data.BookID,
		BookURL:    data.BookURL,
		FSPath:     relativePath,
		ImportedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("create import info: %w", err)
	}

	return nil
}
