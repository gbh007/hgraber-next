package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateArchivePost(
	ctx context.Context,
	req serverapi.APIDeduplicateArchivePostReq,
) (serverapi.APIDeduplicateArchivePostRes, error) {
	data, err := c.deduplicateUseCases.ArchiveEntryPercentage(ctx, req.Data)
	if err != nil {
		return &serverapi.APIDeduplicateArchivePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	result := serverapi.APIDeduplicateArchivePostOKApplicationJSON(
		pkg.Map(data, func(raw core.DeduplicateArchiveResult) serverapi.APIDeduplicateArchivePostOKItem {
			return serverapi.APIDeduplicateArchivePostOKItem{
				BookID:                 raw.TargetBookID,
				BookOriginURL:          apiservercore.OptURL(raw.OriginBookURL),
				EntryPercentage:        raw.EntryPercentage,
				ReverseEntryPercentage: raw.ReverseEntryPercentage,
			}
		}),
	)

	return &result, nil
}
