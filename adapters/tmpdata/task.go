package tmpdata

import "github.com/gbh007/hgraber-next/domain/core"

func (s *Storage) SaveTask(task core.RunnableTask) {
	s.toRun.PushOne(task)
}

func (s *Storage) GetTask() []core.RunnableTask {
	return s.toRun.Pop()
}

func (s *Storage) SaveTaskResult(result *core.TaskResult) {
	s.taskResult.PushOne(result)
}

func (s *Storage) GetTaskResults() []*core.TaskResult {
	return s.taskResult.All()
}
