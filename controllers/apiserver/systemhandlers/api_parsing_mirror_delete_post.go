package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorDeletePost(
	ctx context.Context,
	req *serverapi.APIParsingMirrorDeletePostReq,
) error {
	return c.parseUseCases.DeleteMirror(ctx, req.ID)
}
