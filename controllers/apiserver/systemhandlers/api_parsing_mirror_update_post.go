package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorUpdatePost(ctx context.Context, req *serverapi.APIParsingMirrorUpdatePostReq) (serverapi.APIParsingMirrorUpdatePostRes, error) {
	err := c.parseUseCases.UpdateMirror(ctx, parsing.URLMirror{
		ID:          req.ID,
		Name:        req.Name,
		Prefixes:    req.Prefixes,
		Description: req.Description.Value,
	})
	if err != nil {
		return &serverapi.APIParsingMirrorUpdatePostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIParsingMirrorUpdatePostNoContent{}, nil
}
