package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *SystemHandlersController) APIParsingMirrorListGet(
	ctx context.Context,
) (*serverapi.APIParsingMirrorListGetOK, error) {
	mirrors, err := c.parseUseCases.Mirrors(ctx)
	if err != nil {
		return nil, err
	}

	return &serverapi.APIParsingMirrorListGetOK{
		Mirrors: lo.Map(mirrors, func(mirror parsing.URLMirror, _ int) serverapi.APIParsingMirrorListGetOKMirrorsItem {
			return serverapi.APIParsingMirrorListGetOKMirrorsItem{
				ID:          mirror.ID,
				Name:        mirror.Name,
				Description: apiservercore.OptString(mirror.Description),
				Prefixes:    mirror.Prefixes,
			}
		}),
	}, nil
}
