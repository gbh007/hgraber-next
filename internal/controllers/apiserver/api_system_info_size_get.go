package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APISystemInfoSizeGet(ctx context.Context) (serverAPI.APISystemInfoSizeGetRes, error) {
	info, err := c.webAPIUseCases.SystemSize(ctx)
	if err != nil {
		return &serverAPI.APISystemInfoSizeGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemInfoSizeGetOK{
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
		InvalidFileCount:  int(info.InvalidFileCountByFSSum()),
		DetachedFileCount: int(info.DetachedFileCountByFSSum()),

		DeadHashCount: info.DeadHashCount,

		PagesSize:          info.PageFileSizeByFSSum(),
		PagesSizeFormatted: entities.PrettySize(info.PageFileSizeByFSSum()),
		FilesSize:          info.FileSizeByFSSum(),
		FilesSizeFormatted: entities.PrettySize(info.FileSizeByFSSum()),
	}, nil
}
