package tmpdata

import (
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/pkg"
)

type Storage struct {
	toExport       *pkg.DataQueue[entities.BookFullWithAgent]
	toRun          *pkg.DataQueue[entities.RunnableTask]
	taskResult     *dataList[*entities.TaskResult]
	toValidate     *pkg.DataQueue[uuid.UUID]
	toFileTransfer *pkg.DataQueue[entities.FileTransfer]
}

func New() *Storage {
	return &Storage{
		toExport:       pkg.NewDataQueue[entities.BookFullWithAgent](100),
		toRun:          pkg.NewDataQueue[entities.RunnableTask](10),
		taskResult:     newDataList[*entities.TaskResult](50),
		toValidate:     pkg.NewDataQueue[uuid.UUID](1000),
		toFileTransfer: pkg.NewDataQueue[entities.FileTransfer](1000),
	}
}
