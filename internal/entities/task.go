package entities

import (
	"context"
	"time"
)

type TaskCode byte

const (
	UnknownTaskCode TaskCode = iota
	DeduplicateFilesTaskCode
	RemoveDetachedFilesTaskCode
	RemoveFilesInStoragesMismatchTaskCode
)

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

type TaskResult struct {
	Name        string
	Stages      []*TaskResultStage
	activeStage *TaskResultStage
	Error       string
	StartedAt   time.Time
	EndedAt     time.Time
	Result      string
}

type TaskResultStage struct {
	Name      string
	StartedAt time.Time
	EndedAt   time.Time
	Progress  int64
	Total     int64
	Error     string
	Result    string
}

func (trs *TaskResultStage) Duration() time.Duration {
	return trs.EndedAt.Sub(trs.StartedAt)
}

func (tr *TaskResult) Duration() time.Duration {
	return tr.EndedAt.Sub(tr.StartedAt)
}

func (tr *TaskResult) flush() {
	if tr.activeStage != nil {
		if tr.activeStage.EndedAt.IsZero() {
			tr.activeStage.EndedAt = time.Now()
		}

		tr.activeStage = nil
	}
}

func (tr *TaskResult) StartStage(name string) {
	if tr.StartedAt.IsZero() {
		tr.StartedAt = time.Now()
	}

	tr.flush()
	tr.activeStage = &TaskResultStage{
		Name:      name,
		StartedAt: time.Now(),
	}
	tr.Stages = append(tr.Stages, tr.activeStage)
}

func (tr *TaskResult) EndStage() {
	tr.activeStage.EndedAt = time.Now()
	tr.flush()
}

func (tr *TaskResult) Finish() {
	tr.flush()
	tr.EndedAt = time.Now()
}

func (tr *TaskResult) SetTotal(total int64) {
	if tr.activeStage != nil {
		tr.activeStage.Total = total
	}
}

func (tr *TaskResult) SetProgress(progress int64) {
	if tr.activeStage != nil {
		tr.activeStage.Progress = progress
	}
}

func (tr *TaskResult) IncProgress() {
	if tr.activeStage != nil {
		tr.activeStage.Progress++
	}
}

func (tr *TaskResult) SetError(err error) {
	if err == nil {
		return
	}

	if tr.activeStage != nil {
		tr.activeStage.Error = err.Error()
	} else {
		tr.Error = err.Error()
	}
}

func (tr *TaskResult) SetResult(result string) {
	if result == "" {
		return
	}

	if tr.activeStage != nil {
		tr.activeStage.Result = result
	} else {
		tr.Result = result
	}
}

func (tr *TaskResult) SetName(name string) {
	if name == "" {
		return
	}

	tr.Name = name
}
