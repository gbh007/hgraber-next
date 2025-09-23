package tmpdata

import (
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/pkg"
)

type Storage struct {
	toExport       *pkg.DataQueue[agentmodel.BookToExport]
	toRun          *pkg.DataQueue[systemmodel.RunnableTask]
	taskResult     *dataList[*systemmodel.TaskResult]
	toValidate     *pkg.DataQueue[uuid.UUID]
	toFileTransfer *pkg.DataQueue[fsmodel.FileTransfer]
}

//nolint:mnd // будет исправлено позднее
func New() *Storage {
	return &Storage{
		toExport:       pkg.NewDataQueue[agentmodel.BookToExport](100),
		toRun:          pkg.NewDataQueue[systemmodel.RunnableTask](10),
		taskResult:     newDataList[*systemmodel.TaskResult](50),
		toValidate:     pkg.NewDataQueue[uuid.UUID](1000),
		toFileTransfer: pkg.NewDataQueue[fsmodel.FileTransfer](1000),
	}
}
