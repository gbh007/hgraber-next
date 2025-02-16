package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsUpdatePost(ctx context.Context, req *serverapi.APIFsUpdatePostReq) (serverapi.APIFsUpdatePostRes, error) {
	err := c.fsUseCases.UpdateFileStorage(ctx, fsmodel.FileStorageSystem{
		ID:                  req.ID,
		Name:                req.Name,
		Description:         req.Description.Value,
		AgentID:             req.AgentID.Value,
		Path:                req.Path.Value,
		DownloadPriority:    req.DownloadPriority,
		DeduplicatePriority: req.DeduplicatePriority,
		HighwayEnabled:      req.HighwayEnabled,
		HighwayAddr:         apiservercore.UrlFromOpt(req.HighwayAddr),
	})
	if err != nil {
		return &serverapi.APIFsUpdatePostInternalServerError{
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIFsUpdatePostNoContent{}, nil
}
