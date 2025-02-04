package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APISystemInfoWorkersGet(ctx context.Context) (serverAPI.APISystemInfoWorkersGetRes, error) {
	workers := c.webAPIUseCases.WorkersInfo(ctx)

	return &serverAPI.APISystemInfoWorkersGetOK{
		Workers: pkg.Map(workers, func(w core.SystemWorkerStat) serverAPI.APISystemInfoWorkersGetOKWorkersItem {
			return serverAPI.APISystemInfoWorkersGetOKWorkersItem{
				Name:    w.Name,
				InQueue: w.InQueueCount,
				InWork:  w.InWorkCount,
				Runners: w.RunnersCount,
			}
		}),
	}, nil
}
