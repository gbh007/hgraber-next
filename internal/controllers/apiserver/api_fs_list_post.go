package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIFsListPost(ctx context.Context, req *serverAPI.APIFsListPostReq) (serverAPI.APIFsListPostRes, error) {
	storages, err := c.fsUseCases.FileStoragesWithStatus(ctx, req.IncludeDbFileSize.Value)
	if err != nil {
		return &serverAPI.APIFsListPostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsListPostOK{
		FileSystems: pkg.Map(storages, func(raw entities.FSWithStatus) serverAPI.APIFsListPostOKFileSystemsItem {
			dbInfo := serverAPI.OptAPIFsListPostOKFileSystemsItemDbFilesInfo{}

			if raw.DBFile != nil {
				dbInfo = serverAPI.NewOptAPIFsListPostOKFileSystemsItemDbFilesInfo(serverAPI.APIFsListPostOKFileSystemsItemDbFilesInfo{
					Count:         raw.DBFile.Count,
					Size:          raw.DBFile.Size,
					SizeFormatted: entities.PrettySize(raw.DBFile.Size),
				})
			}

			return serverAPI.APIFsListPostOKFileSystemsItem{
				Info:        convertFileSystemInfoToAPI(raw.Info),
				IsLegacy:    raw.IsLegacy,
				DbFilesInfo: dbInfo,
			}
		}),
	}, nil
}
