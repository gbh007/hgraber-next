package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APISystemInfoSizeGet(ctx context.Context) (serverapi.APISystemInfoSizeGetRes, error) {
	info, err := c.webAPIUseCases.SystemSize(ctx)
	if err != nil {
		return &serverapi.APISystemInfoSizeGetInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APISystemInfoSizeGetOK{
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
