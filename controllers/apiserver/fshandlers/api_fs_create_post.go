package fshandlers

import (
	"context"

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
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIFsCreatePostOK{
		ID: id,
	}, nil
}
