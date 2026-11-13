package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/systemmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *SystemHandlersController) APISystemInfoWorkersGet(
	ctx context.Context,
) (*serverapi.APISystemInfoWorkersGetOK, error) {
	workers := c.systemUseCases.WorkersInfo(ctx)

	return &serverapi.APISystemInfoWorkersGetOK{
		Workers: lo.Map(workers, func(w systemmodel.SystemWorkerStat, _ int) serverapi.APISystemInfoWorkersGetOKWorkersItem {
			return serverapi.APISystemInfoWorkersGetOKWorkersItem{
				Name:    w.Name,
				InQueue: w.InQueueCount,
				InWork:  w.InWorkCount,
				Runners: w.RunnersCount,
			}
		}),
	}, nil
}
