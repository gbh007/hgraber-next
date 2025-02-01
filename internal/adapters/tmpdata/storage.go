package tmpdata

import (
	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type Storage struct {
	toExport       *dataQueue[entities.BookFullWithAgent]
	toRun          *dataQueue[entities.RunnableTask]
	taskResult     *dataList[*entities.TaskResult]
	toValidate     *dataQueue[uuid.UUID]
	toFileTransfer *dataQueue[entities.FileTransfer]
}

func New() *Storage {
	return &Storage{
		toExport:       newDataQueue[entities.BookFullWithAgent](100),
		toRun:          newDataQueue[entities.RunnableTask](10),
		taskResult:     newDataList[*entities.TaskResult](50),
		toValidate:     newDataQueue[uuid.UUID](1000),
		toFileTransfer: newDataQueue[entities.FileTransfer](1000),
	}
}
