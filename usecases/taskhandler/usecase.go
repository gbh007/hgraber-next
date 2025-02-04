package taskhandler

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	SaveTask(task core.RunnableTask)
	GetTaskResults() []*core.TaskResult
}

type deduplicator interface {
	DeduplicateFiles(ctx context.Context) (core.RunnableTask, error)
	FillDeadHashes(ctx context.Context, withRemoveDeletedPages bool) (core.RunnableTask, error)
}

type cleanuper interface {
	RemoveDetachedFiles(ctx context.Context) (core.RunnableTask, error)
	RemoveFilesInStoragesMismatch(ctx context.Context, fsID uuid.UUID) (core.RunnableTask, error)
	CleanDeletedPages(ctx context.Context) (core.RunnableTask, error)
	CleanDeletedRebuilds(ctx context.Context) (core.RunnableTask, error)
}

type UseCase struct {
	logger *slog.Logger

	storage      storage
	deduplicator deduplicator
	cleanuper    cleanuper
}

func New(
	logger *slog.Logger,
	storage storage,
	deduplicator deduplicator,
	cleanuper cleanuper,
) *UseCase {
	return &UseCase{
		logger:       logger,
		storage:      storage,
		deduplicator: deduplicator,
		cleanuper:    cleanuper,
	}
}

func (uc *UseCase) RunTask(ctx context.Context, code core.TaskCode) error {
	var (
		task core.RunnableTask
		err  error
	)

	switch code {
	case core.DeduplicateFilesTaskCode:
		task, err = uc.deduplicator.DeduplicateFiles(ctx)
	case core.RemoveDetachedFilesTaskCode:
		task, err = uc.cleanuper.RemoveDetachedFiles(ctx)
	case core.FillDeadHashesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, false)
	case core.FillDeadHashesAndRemoveDeletedPagesTaskCode:
		task, err = uc.deduplicator.FillDeadHashes(ctx, true)
	case core.CleanDeletedPagesTaskCode:
		task, err = uc.cleanuper.CleanDeletedPages(ctx)
	case core.CleanDeletedRebuildsTaskCode:
		task, err = uc.cleanuper.CleanDeletedRebuilds(ctx)
	}

	if err != nil {
		return err
	}

	if task != nil {
		uc.storage.SaveTask(task)
	}

	return nil
}

func (uc *UseCase) TaskResults(ctx context.Context) ([]*core.TaskResult, error) {
	return uc.storage.GetTaskResults(), nil
}

func (uc *UseCase) RemoveFilesInFSMismatch(ctx context.Context, fsID uuid.UUID) error {
	task, err := uc.cleanuper.RemoveFilesInStoragesMismatch(ctx, fsID)
	if err != nil {
		return err
	}

	uc.storage.SaveTask(task)

	return nil
}
