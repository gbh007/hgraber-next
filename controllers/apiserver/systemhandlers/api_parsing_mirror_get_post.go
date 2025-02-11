package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorGetPost(ctx context.Context, req *serverapi.APIParsingMirrorGetPostReq) (serverapi.APIParsingMirrorGetPostRes, error) {
	mirror, err := c.parseUseCases.Mirror(ctx, req.ID)
	if err != nil {
		return &serverapi.APIParsingMirrorGetPostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIParsingMirrorGetPostOK{
		ID:          mirror.ID,
		Name:        mirror.Name,
		Description: apiservercore.OptString(mirror.Description),
		Prefixes:    mirror.Prefixes,
	}, nil
}
