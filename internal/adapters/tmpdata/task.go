package tmpdata

import "github.com/gbh007/hgraber-next/internal/entities"

func (s *Storage) SaveTask(task entities.RunnableTask) {
	s.toRun.PushOne(task)
}

func (s *Storage) GetTask() []entities.RunnableTask {
	return s.toRun.Pop()
}

func (s *Storage) SaveTaskResult(result *entities.TaskResult) {
	s.taskResult.PushOne(result)
}

func (s *Storage) GetTaskResults() []*entities.TaskResult {
	return s.taskResult.All()
}
