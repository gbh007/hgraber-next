package deduplicatehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *DeduplicateHandlersController) APIDeduplicateArchivePost(
	ctx context.Context,
	req serverapi.APIDeduplicateArchivePostReq,
) ([]serverapi.APIDeduplicateArchivePostOKItem, error) {
	data, err := c.deduplicateUseCases.ArchiveEntryPercentage(ctx, req.Data)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return lo.Map(data, func(raw core.DeduplicateArchiveResult, _ int) serverapi.APIDeduplicateArchivePostOKItem {
		return serverapi.APIDeduplicateArchivePostOKItem{
			BookID:                 raw.TargetBookID,
			BookOriginURL:          apiservercore.OptURL(raw.OriginBookURL),
			EntryPercentage:        raw.EntryPercentage,
			ReverseEntryPercentage: raw.ReverseEntryPercentage,
		}
	}), nil
}
