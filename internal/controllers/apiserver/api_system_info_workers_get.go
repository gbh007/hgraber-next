package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemInfoWorkersGet(ctx context.Context) (serverAPI.APISystemInfoWorkersGetRes, error) {
	workers := c.webAPIUseCases.WorkersInfo(ctx)

	return &serverAPI.APISystemInfoWorkersGetOK{
		Workers: pkg.Map(workers, func(w entities.SystemWorkerStat) serverAPI.APISystemInfoWorkersGetOKWorkersItem {
			return serverAPI.APISystemInfoWorkersGetOKWorkersItem{
				Name:    w.Name,
				InQueue: w.InQueueCount,
				InWork:  w.InWorkCount,
				Runners: w.RunnersCount,
			}
		}),
	}, nil
}
