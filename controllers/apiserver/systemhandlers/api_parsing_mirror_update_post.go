package systemhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APIParsingMirrorUpdatePost(
	ctx context.Context,
	req *serverapi.APIParsingMirrorUpdatePostReq,
) error {
	err := c.parseUseCases.UpdateMirror(ctx, parsing.URLMirror{
		ID:          req.ID,
		Name:        req.Name,
		Prefixes:    req.Prefixes,
		Description: req.Description.Value,
	})
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
