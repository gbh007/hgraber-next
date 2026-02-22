package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorUpdatePost(
	ctx context.Context,
	req *serverapi.APIParsingMirrorUpdatePostReq,
) error {
	return c.parseUseCases.UpdateMirror(ctx, parsing.URLMirror{
		ID:          req.ID,
		Name:        req.Name,
		Prefixes:    req.Prefixes,
		Description: req.Description.Value,
	})
}
