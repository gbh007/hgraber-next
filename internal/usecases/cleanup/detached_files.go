package cleanup

import (
	"context"
	"fmt"
	"log/slog"
)

func (uc *UseCase) RemoveDetachedFiles(ctx context.Context) (count int, size int64, err error) {
	files, err := uc.storage.DetachedFiles(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("get detached files from storage: %w", err)
	}

	for _, file := range files {
		if file.Size == 0 {
			uc.logger.WarnContext(
				ctx, "empty file size",
				slog.String("id", file.ID.String()),
			)

			continue
		}

		err = uc.storage.DeleteFile(ctx, file.ID)
		if err != nil {
			return 0, 0, fmt.Errorf("delete file (%s) from storage: %w", file.ID.String(), err)
		}

		err = uc.fileStorage.Delete(ctx, file.ID)
		if err != nil {
			return 0, 0, fmt.Errorf("delete file (%s) from file-storage: %w", file.ID.String(), err)
		}

		count++

		size += file.Size
	}

	return count, size, nil
}
