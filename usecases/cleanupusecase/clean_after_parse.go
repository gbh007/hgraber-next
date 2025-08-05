package cleanupusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

func (uc *UseCase) CleanAfterParse(_ context.Context) (systemmodel.RunnableTask, error) {
	return systemmodel.RunnableTaskFunction(func(ctx context.Context, taskResult systemmodel.TaskResultWriter) {
		defer taskResult.Finish()

		taskResult.SetName("CleanAfterParse")

		deduplicateFileCount, deduplicateFileSize, err := uc.deduplicateFiles(ctx, taskResult)
		if err != nil {
			return // В результат выполнения данные уже проставлены
		}

		detachedFileCount, detachedFileSize, err := uc.removeDetachedFiles(ctx, taskResult)
		if err != nil {
			return // В результат выполнения данные уже проставлены
		}

		taskResult.SetResult(fmt.Sprintf(
			"Дедуплицированные файлы = количество: %d байт: %d размер: %s\n"+
				"Удаленные файлы = количество: %d байт: %d размер: %s\n",
			deduplicateFileCount,
			deduplicateFileSize,
			core.PrettySize(deduplicateFileSize),
			detachedFileCount,
			detachedFileSize,
			core.PrettySize(detachedFileSize),
		))
	}), nil
}
