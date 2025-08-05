package systemusecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

func (uc *UseCase) RunTask(ctx context.Context, code systemmodel.TaskCode) error {
	var (
		task systemmodel.RunnableTask
		err  error
	)

	switch code {
	case systemmodel.DeduplicateFilesTaskCode:
		task, err = uc.deduplicator.DeduplicateFiles(ctx)
	case systemmodel.RemoveDetachedFilesTaskCode:
		task, err = uc.cleanuper.RemoveDetachedFiles(ctx)
	case systemmodel.FillDeadHashesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, false)
	case systemmodel.FillDeadHashesAndRemoveDeletedPagesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, true)
	case systemmodel.CleanDeletedPagesTaskCode:
		task, err = uc.cleanuper.CleanDeletedPages(ctx)
	case systemmodel.CleanDeletedRebuildsTaskCode:
		task, err = uc.cleanuper.CleanDeletedRebuilds(ctx)
	case systemmodel.RemapAttributesTaskCode:
		task, err = uc.attributeRemaper.RemapBooks(ctx)
	case systemmodel.CleanAfterRebuildTaskCode:
		task, err = uc.cleanuper.CleanAfterRebuild(ctx)
	}

	if err != nil {
		return err
	}

	if task != nil {
		uc.tmpStorage.SaveTask(task)
	}

	return nil
}

func (uc *UseCase) TaskResults(ctx context.Context) ([]*systemmodel.TaskResult, error) {
	return uc.tmpStorage.GetTaskResults(), nil
}

func (uc *UseCase) RemoveFilesInFSMismatch(ctx context.Context, fsID uuid.UUID) error {
	task, err := uc.cleanuper.RemoveFilesInStoragesMismatch(ctx, fsID)
	if err != nil {
		return err
	}

	uc.tmpStorage.SaveTask(task)

	return nil
}
