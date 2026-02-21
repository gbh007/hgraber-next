package fshandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsCreatePost(
	ctx context.Context,
	req *serverapi.APIFsCreatePostReq,
) (*serverapi.APIFsCreatePostOK, error) {
	id, err := c.fsUseCases.NewFileStorage(ctx, fsmodel.FileStorageSystem{
		Name:                req.Name,
		Description:         req.Description.Value,
		AgentID:             req.AgentID.Value,
		Path:                req.Path.Value,
		DownloadPriority:    req.DownloadPriority,
		DeduplicatePriority: req.DeduplicatePriority,
		HighwayEnabled:      req.HighwayEnabled,
		HighwayAddr:         apiservercore.URLFromOpt(req.HighwayAddr),
	})
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.FSUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APIFsCreatePostOK{
		ID: id,
	}, nil
}
