package agentcacheapi

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2FsArchivePost(
	ctx context.Context,
	req agentapi.APIV2FsArchivePostReq,
	params agentapi.APIV2FsArchivePostParams,
) error {
	_, err := c.exportUseCases.ImportArchive(ctx, req.Data, true, false)
	if err != nil {
		return err //nolint:wrapcheck // будет исправлено позднее
	}

	return nil
}
