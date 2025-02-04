package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIFsListPost(ctx context.Context, req *serverAPI.APIFsListPostReq) (serverAPI.APIFsListPostRes, error) {
	storages, err := c.fsUseCases.FileStoragesWithStatus(
		ctx,
		req.IncludeDbFileSize.Value,
		req.IncludeAvailableSize.Value,
	)
	if err != nil {
		return &serverAPI.APIFsListPostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsListPostOK{
		FileSystems: pkg.Map(storages, func(raw entities.FSWithStatus) serverAPI.APIFsListPostOKFileSystemsItem {
			return serverAPI.APIFsListPostOKFileSystemsItem{
				Info:                convertFileSystemInfoToAPI(raw.Info),
				IsLegacy:            raw.IsLegacy,
				DbFilesInfo:         convertFSDBFilesInfoToAPI(raw.DBFile),
				DbInvalidFilesInfo:  convertFSDBFilesInfoToAPI(raw.DBInvalidFile),
				DbDetachedFilesInfo: convertFSDBFilesInfoToAPI(raw.DBDetachedFile),
				AvailableSize: serverAPI.OptInt64{
					Value: raw.AvailableSize,
					Set:   raw.AvailableSize > 0,
				},
				AvailableSizeFormatted: serverAPI.OptString{
					Value: entities.PrettySize(raw.AvailableSize),
					Set:   raw.AvailableSize > 0,
				},
			}
		}),
	}, nil
}
