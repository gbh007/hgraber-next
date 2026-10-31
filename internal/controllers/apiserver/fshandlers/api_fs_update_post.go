package fshandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/fsmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFsUpdatePost(
	ctx context.Context,
	req *serverapi.APIFsUpdatePostReq,
) error {
	return c.fsUseCases.UpdateFileStorage(ctx, fsmodel.FileStorageSystem{
		ID:                  req.ID,
		Name:                req.Name,
		Description:         req.Description.Value,
		AgentID:             req.AgentID.Value,
		Path:                req.Path.Value,
		DownloadPriority:    req.DownloadPriority,
		DeduplicatePriority: req.DeduplicatePriority,
		HighwayEnabled:      req.HighwayEnabled,
		HighwayAddr:         apiservercore.URLFromOpt(req.HighwayAddr),
	})
}
