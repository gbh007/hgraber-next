package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *FSHandlersController) APIFsListPost(ctx context.Context, req *serverAPI.APIFsListPostReq) (serverAPI.APIFsListPostRes, error) {
	storages, err := c.fsUseCases.FileStoragesWithStatus(
		ctx,
		req.IncludeDbFileSize.Value,
		req.IncludeAvailableSize.Value,
	)
	if err != nil {
		return &serverAPI.APIFsListPostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsListPostOK{
		FileSystems: pkg.Map(storages, func(raw core.FSWithStatus) serverAPI.APIFsListPostOKFileSystemsItem {
			return serverAPI.APIFsListPostOKFileSystemsItem{
				Info:                apiservercore.ConvertFileSystemInfoToAPI(raw.Info),
				IsLegacy:            raw.IsLegacy,
				DbFilesInfo:         apiservercore.ConvertFSDBFilesInfoToAPI(raw.DBFile),
				DbInvalidFilesInfo:  apiservercore.ConvertFSDBFilesInfoToAPI(raw.DBInvalidFile),
				DbDetachedFilesInfo: apiservercore.ConvertFSDBFilesInfoToAPI(raw.DBDetachedFile),
				AvailableSize: serverAPI.OptInt64{
					Value: raw.AvailableSize,
					Set:   raw.AvailableSize > 0,
				},
				AvailableSizeFormatted: serverAPI.OptString{
					Value: core.PrettySize(raw.AvailableSize),
					Set:   raw.AvailableSize > 0,
				},
			}
		}),
	}, nil
}
