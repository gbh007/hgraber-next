package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APISystemInfoGet(ctx context.Context) (server.APISystemInfoGetRes, error) {
	info, err := c.webAPIUseCases.SystemInfo(ctx)
	if err != nil {
		return &server.APISystemInfoGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.SystemInfo{
		Count:                info.BookCount,
		NotLoadCount:         info.BookUnparsedCount,
		PageCount:            info.PageCount,
		NotLoadPageCount:     info.PageUnloadedCount,
		PageWithoutBodyCount: info.PageWithoutBodyCount,
		PagesSize:            info.PageFileSize,
		PagesSizeFormatted:   entities.PrettySize(info.PageFileSize),
		Monitor: server.NewOptSystemInfoMonitor(server.SystemInfoMonitor{
			Workers: pkg.Map(info.Workers, func(w entities.SystemWorkerStat) server.SystemInfoMonitorWorkersItem {
				return server.SystemInfoMonitorWorkersItem{
					Name:    w.Name,
					InQueue: w.InQueueCount,
					InWork:  w.InWorkCount,
					Runners: w.RunnersCount,
				}
			}),
		}),
	}, nil
}
