package tmpdata

import (
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

type Storage struct {
	toExport       *pkg.DataQueue[core.BookFullWithAgent]
	toRun          *pkg.DataQueue[core.RunnableTask]
	taskResult     *dataList[*core.TaskResult]
	toValidate     *pkg.DataQueue[uuid.UUID]
	toFileTransfer *pkg.DataQueue[core.FileTransfer]
}

func New() *Storage {
	return &Storage{
		toExport:       pkg.NewDataQueue[core.BookFullWithAgent](100),
		toRun:          pkg.NewDataQueue[core.RunnableTask](10),
		taskResult:     newDataList[*core.TaskResult](50),
		toValidate:     pkg.NewDataQueue[uuid.UUID](1000),
		toFileTransfer: pkg.NewDataQueue[core.FileTransfer](1000),
	}
}
