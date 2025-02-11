package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorCreatePost(ctx context.Context, req *serverapi.APIParsingMirrorCreatePostReq) (serverapi.APIParsingMirrorCreatePostRes, error) {
	err := c.parseUseCases.NewMirror(ctx, parsing.URLMirror{
		Name:        req.Name,
		Prefixes:    req.Prefixes,
		Description: req.Description.Value,
	})
	if err != nil {
		return &serverapi.APIParsingMirrorCreatePostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIParsingMirrorCreatePostNoContent{}, nil
}
