package systemhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APIParsingMirrorListGet(
	ctx context.Context,
) (*serverapi.APIParsingMirrorListGetOK, error) {
	mirrors, err := c.parseUseCases.Mirrors(ctx)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APIParsingMirrorListGetOK{
		Mirrors: pkg.Map(mirrors, func(mirror parsing.URLMirror) serverapi.APIParsingMirrorListGetOKMirrorsItem {
			return serverapi.APIParsingMirrorListGetOKMirrorsItem{
				ID:          mirror.ID,
				Name:        mirror.Name,
				Description: apiservercore.OptString(mirror.Description),
				Prefixes:    mirror.Prefixes,
			}
		}),
	}, nil
}
