package tmpdata

import "hgnext/internal/entities"

type Storage struct {
	toExport   *dataQueue[entities.BookFullWithAgent]
	toRun      *dataQueue[entities.RunnableTask]
	taskResult *dataList[*entities.TaskResult]
}

func New() *Storage {
	return &Storage{
		toExport:   newDataQueue[entities.BookFullWithAgent](100),
		toRun:      newDataQueue[entities.RunnableTask](10),
		taskResult: newDataList[*entities.TaskResult](50),
	}
}
