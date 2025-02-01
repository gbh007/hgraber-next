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
		Count:           info.BookCount,
		DownloadedCount: info.DownloadedBookCount,
		VerifiedCount:   info.VerifiedBookCount,
		RebuildedCount:  info.RebuildedBookCount,
		NotLoadCount:    info.BookUnparsedCount,
		DeletedCount:    info.DeletedBookCount,

		PageCount:            info.PageCount,
		NotLoadPageCount:     info.PageUnloadedCount,
		PageWithoutBodyCount: info.PageWithoutBodyCount,
		DeletedPageCount:     info.DeletedPageCount,

		FileCount:         int(info.FileCountByFSSum()),
		UnhashedFileCount: int(info.UnhashedFileCountByFSSum()),
		DeadHashCount:     info.DeadHashCount,

		PagesSize:          info.PageFileSizeByFSSum(),
		PagesSizeFormatted: entities.PrettySize(info.PageFileSizeByFSSum()),
		FilesSize:          info.FileSizeByFSSum(),
		FilesSizeFormatted: entities.PrettySize(info.FileSizeByFSSum()),

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
