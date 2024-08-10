package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemInfoGet(ctx context.Context) (serverAPI.APISystemInfoGetRes, error) {
	info, err := c.webAPIUseCases.SystemInfo(ctx)
	if err != nil {
		return &serverAPI.APISystemInfoGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.SystemInfo{
		Count:                info.BookCount,
		NotLoadCount:         info.BookUnparsedCount,
		PageCount:            info.PageCount,
		NotLoadPageCount:     info.PageUnloadedCount,
		PageWithoutBodyCount: info.PageWithoutBodyCount,
		PagesSize:            info.PageFileSize,
		PagesSizeFormatted:   entities.PrettySize(info.PageFileSize),
		FilesSize:            info.FileSize,
		FilesSizeFormatted:   entities.PrettySize(info.FileSize),
		Monitor: serverAPI.NewOptSystemInfoMonitor(serverAPI.SystemInfoMonitor{
			Workers: pkg.Map(info.Workers, func(w entities.SystemWorkerStat) serverAPI.SystemInfoMonitorWorkersItem {
				return serverAPI.SystemInfoMonitorWorkersItem{
					Name:    w.Name,
					InQueue: w.InQueueCount,
					InWork:  w.InWorkCount,
					Runners: w.RunnersCount,
				}
			}),
		}),
	}, nil
}
