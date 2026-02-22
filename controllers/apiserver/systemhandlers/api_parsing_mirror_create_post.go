package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorCreatePost(
	ctx context.Context,
	req *serverapi.APIParsingMirrorCreatePostReq,
) error {
	return c.parseUseCases.NewMirror(ctx, parsing.URLMirror{
		Name:        req.Name,
		Prefixes:    req.Prefixes,
		Description: req.Description.Value,
	})
}
