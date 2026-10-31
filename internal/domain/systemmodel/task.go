package systemmodel

import (
	"context"
)

const (
	UnknownTaskCode TaskCode = iota
	DeduplicateFilesTaskCode
	RemoveDetachedFilesTaskCode
	FillDeadHashesTaskCode
	FillDeadHashesAndRemoveDeletedPagesTaskCode
	CleanDeletedPagesTaskCode
	CleanDeletedRebuildsTaskCode
	RemapAttributesTaskCode
	CleanAfterRebuildTaskCode
	CleanAfterParseTaskCode
)

type TaskCode byte

type TaskResultWriter interface {
	EndStage()
	Finish()
	SetError(err error)
	SetResult(result string)
	SetName(name string)
	SetProgress(progress int64)
	IncProgress()
	SetTotal(total int64)
	StartStage(name string)
}

type RunnableTask interface {
	Run(ctx context.Context, taskResult TaskResultWriter)
}

type RunnableTaskFunction func(ctx context.Context, taskResult TaskResultWriter)

func (f RunnableTaskFunction) Run(ctx context.Context, taskResult TaskResultWriter) {
	f(ctx, taskResult)
}
