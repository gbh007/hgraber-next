package localFiles

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (s *Storage) State(ctx context.Context, includeFileIDs, includeFileSizes bool) (fsmodel.FSState, error) {
	var state fsmodel.FSState

	if includeFileIDs || includeFileSizes {
		entries, err := os.ReadDir(s.fsPath)
		if err != nil {
			return fsmodel.FSState{}, fmt.Errorf("local fs: scan dir: %w", err)
		}

		if includeFileIDs {
			state.FileIDs = make([]uuid.UUID, 0, len(entries))
		}

		if includeFileSizes {
			state.Files = make([]fsmodel.FSStateFile, 0, len(entries))
		}

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

			state.TotalFileCount++

			if includeFileSizes {
				stat, err := os.Stat(path.Join(s.fsPath, e.Name()))
				if err != nil {
					return fsmodel.FSState{}, fmt.Errorf("get file (%s) stat: %w", e.Name(), err)
				}

				state.Files = append(state.Files, fsmodel.FSStateFile{
					ID:        id,
					Size:      stat.Size(),
					CreatedAt: stat.ModTime(),
				})

				state.TotalFileSize += stat.Size()
			}

			if includeFileIDs {
				state.FileIDs = append(state.FileIDs, id)
			}
		}
	}

	state.AvailableSize = getAvailableSize(s.fsPath)

	return state, nil
}
