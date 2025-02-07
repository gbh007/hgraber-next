package tmpdata

import "github.com/gbh007/hgraber-next/domain/systemmodel"

func (s *Storage) SaveTask(task systemmodel.RunnableTask) {
	s.toRun.PushOne(task)
}

func (s *Storage) GetTask() []systemmodel.RunnableTask {
	return s.toRun.Pop()
}

func (s *Storage) SaveTaskResult(result *systemmodel.TaskResult) {
	s.taskResult.PushOne(result)
}

func (s *Storage) GetTaskResults() []*systemmodel.TaskResult {
	return s.taskResult.All()
}
