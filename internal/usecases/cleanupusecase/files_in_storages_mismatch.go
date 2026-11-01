package cleanupusecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/domain/fsmodel"
	"github.com/gbh007/hgraber-next/internal/domain/systemmodel"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

func (uc *UseCase) RemoveFilesInStoragesMismatch(_ context.Context, fsID uuid.UUID) (systemmodel.RunnableTask, error) {
	return systemmodel.RunnableTaskFunction(func(ctx context.Context, taskResult systemmodel.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("RemoveFilesInStoragesMismatch")

		taskResult.StartStage("search file ids in fs")

		fsState, err := uc.fileStorage.State(ctx, true, false, fsID)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		fileIDs := fsState.FileIDs

		taskResult.SetTotal(int64(len(fileIDs)))
		taskResult.SetProgress(int64(len(fileIDs)))
		taskResult.EndStage()

		taskResult.StartStage("search file ids in storage")

		storageIDs, err := uc.storage.FileIDsByFilter(ctx, fsmodel.FileFilter{FSID: &fsID})
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.SetTotal(int64(len(storageIDs)))
		taskResult.SetProgress(int64(len(storageIDs)))
		taskResult.EndStage()

		taskResult.StartStage("transform ids")
		taskResult.SetTotal(int64(len(storageIDs)*2 + len(fileIDs)*2))

		fileNotInDB := pkg.SliceToMap(fileIDs, func(id uuid.UUID) (uuid.UUID, struct{}) {
			taskResult.IncProgress()

			return id, struct{}{}
		})
		fileNotInFS := pkg.SliceToMap(storageIDs, func(id uuid.UUID) (uuid.UUID, struct{}) {
			taskResult.IncProgress()

			return id, struct{}{}
		})

		for _, id := range fileIDs {
			taskResult.IncProgress()
			delete(fileNotInFS, id)
		}

		for _, id := range storageIDs {
			taskResult.IncProgress()
			delete(fileNotInDB, id)
		}

		taskResult.EndStage()

		uc.logger.DebugContext(
			ctx, "RemoveFilesInStoragesMismatch",
			slog.Int("file_not_in_fs_count", len(fileNotInFS)),
			slog.Int("file_not_in_db_count", len(fileNotInDB)),
			slog.Any("file_not_in_fs", fileNotInFS),
			slog.Any("file_not_in_db", fileNotInDB),
		)

		taskResult.StartStage("remove files from fs")
		taskResult.SetTotal(int64(len(fileNotInDB)))

		for id := range fileNotInDB {
			taskResult.IncProgress()

			err = uc.fileStorage.Delete(ctx, id, &fsID)
			if err != nil {
				taskResult.SetError(err)

				return
			}
		}

		taskResult.EndStage()

		taskResult.StartStage("remove files from storage")
		taskResult.SetTotal(int64(len(fileNotInFS)))

		for id := range fileNotInFS {
			taskResult.IncProgress()

			err = uc.storage.DeleteFile(ctx, id)
			if err != nil {
				taskResult.SetError(err)

				return
			}
		}

		taskResult.EndStage()

		taskResult.SetResult(fmt.Sprintf(
			"fs id: %s\nremove from fs: %d remove from db: %d",
			fsID.String(), len(fileNotInDB), len(fileNotInFS),
		))
	}), nil
}
