package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APISystemInfoWorkersGet(ctx context.Context) (serverapi.APISystemInfoWorkersGetRes, error) {
	workers := c.systemUseCases.WorkersInfo(ctx)

	return &serverapi.APISystemInfoWorkersGetOK{
		Workers: pkg.Map(workers, func(w systemmodel.SystemWorkerStat) serverapi.APISystemInfoWorkersGetOKWorkersItem {
			return serverapi.APISystemInfoWorkersGetOKWorkersItem{
				Name:    w.Name,
				InQueue: w.InQueueCount,
				InWork:  w.InWorkCount,
				Runners: w.RunnersCount,
			}
		}),
	}, nil
}
