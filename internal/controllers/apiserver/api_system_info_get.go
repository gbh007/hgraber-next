package apiserver

import (
	"context"
	"strconv"

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
		Count:              info.BookCount,
		NotLoadCount:       info.BookUnparsedCount,
		PageCount:          info.PageCount,
		NotLoadPageCount:   info.PageUnloadedCount,
		PagesSize:          info.PageFileSize,
		PagesSizeFormatted: prettySize(info.PageFileSize),
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

func prettySize(raw int64) string {
	if raw < 1 {
		return "? б"
	}

	var div, mod int64

	const divider = 1024

	div = raw
	step := 0

	for div/divider > 0 {
		step++

		mod = div % divider
		div = div / divider
	}

	return strconv.FormatInt(div, 10) + "." + strconv.FormatInt(mod*10/1024, 10) + " " + sizeUnitFromStep(step)
}

func sizeUnitFromStep(step int) string {
	switch step {
	case 0:
		return "б"
	case 1:
		return "Кб"
	case 2:
		return "Мб"
	case 3:
		return "Гб"
	case 4:
		return "Тб"
	default:
		return "??"
	}
}
