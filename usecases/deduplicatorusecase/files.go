package deduplicatorusecase

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) DeduplicateFiles(_ context.Context) (systemmodel.RunnableTask, error) {
	return systemmodel.RunnableTaskFunction(func(ctx context.Context, taskResult systemmodel.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("DeduplicateFiles")

		ctx, span := uc.tracer.Start(ctx, "DeduplicateFiles")
		defer span.End()

		taskResult.StartStage("get fs from db")
		span.AddEvent("get fs from db", trace.WithTimestamp(time.Now()))

		storages, err := uc.storage.FileStorages(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		storageMap := pkg.SliceToMap(storages, func(s fsmodel.FileStorageSystem) (uuid.UUID, fsmodel.FileStorageSystem) {
			return s.ID, s
		})

		taskResult.StartStage("get duplicates from storage")
		span.AddEvent("get duplicates from storage", trace.WithTimestamp(time.Now()))

		files, err := uc.storage.DuplicatedFiles(ctx)
		if err != nil {
			taskResult.SetError(err)

			return
		}

		taskResult.SetTotal(int64(len(files)))
		taskResult.SetProgress(int64(len(files)))
		taskResult.EndStage()

		taskResult.StartStage("transform data")
		taskResult.SetTotal(int64(len(files)))
		span.AddEvent("transform data", trace.WithTimestamp(time.Now()))

		fileMap := make(map[core.FileHash][]core.File)

		for _, file := range files {
			taskResult.IncProgress()

			k := file.Hash()

			fileMap[k] = append(fileMap[k], file)
		}

		taskResult.EndStage()

		taskResult.StartStage("handle duplicates")
		taskResult.SetTotal(int64(len(fileMap)))
		span.AddEvent("handle duplicates", trace.WithTimestamp(time.Now()))

		var (
			count int
			size  int64
		)

		for k, files := range fileMap {
			taskResult.IncProgress()

			if k.Size == 0 {
				uc.logger.WarnContext(
					ctx, "empty file size",
					slog.Any("ids", files),
				)

				continue
			}

			if len(files) < 2 {
				uc.logger.WarnContext(
					ctx, "invalid deduplicate ids len",
					slog.Any("ids", files),
				)

				continue
			}

			// Оставляем файл с наивысшим приоритетом его ФС для сохранения
			slices.SortStableFunc(files, func(a, b core.File) int {
				return storageMap[b.FSID].DeduplicatePriority - storageMap[a.FSID].DeduplicatePriority
			})

			newID := files[0].ID
			filesToRemove := files[1:]

			for _, file := range filesToRemove {
				err = uc.storage.ReplaceFile(ctx, file.ID, newID)
				if err != nil {
					taskResult.SetError(fmt.Errorf("replace id in storage: %w", err))

					return
				}

				uc.logger.InfoContext(
					ctx, "replaced file",
					slog.String("old", file.ID.String()),
					slog.String("new", newID.String()),
				)
			}

			size += k.Size * int64(len(filesToRemove))
			count += len(filesToRemove)
		}

		taskResult.EndStage()

		taskResult.SetResult(fmt.Sprintf(
			"count: %d size: %d human size: %s",
			count, size, core.PrettySize(size),
		))
	}), nil
}
