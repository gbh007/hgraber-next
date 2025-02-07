package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *FSHandlersController) APIFsListPost(ctx context.Context, req *serverapi.APIFsListPostReq) (serverapi.APIFsListPostRes, error) {
	storages, err := c.fsUseCases.FileStoragesWithStatus(
		ctx,
		req.IncludeDbFileSize.Value,
		req.IncludeAvailableSize.Value,
	)
	if err != nil {
		return &serverapi.APIFsListPostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsListPostOK{
		FileSystems: pkg.Map(storages, func(raw fsmodel.FSWithStatus) serverapi.APIFsListPostOKFileSystemsItem {
			return serverapi.APIFsListPostOKFileSystemsItem{
				Info:                apiservercore.ConvertFileSystemInfoToAPI(raw.Info),
				IsLegacy:            raw.IsLegacy,
				DbFilesInfo:         apiservercore.ConvertFSDBFilesInfoToAPI(raw.DBFile),
				DbInvalidFilesInfo:  apiservercore.ConvertFSDBFilesInfoToAPI(raw.DBInvalidFile),
				DbDetachedFilesInfo: apiservercore.ConvertFSDBFilesInfoToAPI(raw.DBDetachedFile),
				AvailableSize: serverapi.OptInt64{
					Value: raw.AvailableSize,
					Set:   raw.AvailableSize > 0,
				},
				AvailableSizeFormatted: serverapi.OptString{
					Value: core.PrettySize(raw.AvailableSize),
					Set:   raw.AvailableSize > 0,
				},
			}
		}),
	}, nil
}
