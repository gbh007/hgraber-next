package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *SystemHandlersController) APISystemInfoSizeGet(ctx context.Context) (serverAPI.APISystemInfoSizeGetRes, error) {
	info, err := c.webAPIUseCases.SystemSize(ctx)
	if err != nil {
		return &serverAPI.APISystemInfoSizeGetInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
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
		PagesSizeFormatted: core.PrettySize(info.PageFileSizeByFSSum()),
		FilesSize:          info.FileSizeByFSSum(),
		FilesSizeFormatted: core.PrettySize(info.FileSizeByFSSum()),
	}, nil
}
