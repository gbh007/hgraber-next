package deduplicatehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *DeduplicateHandlersController) APIDeduplicateArchivePost(
	ctx context.Context,
	req serverapi.APIDeduplicateArchivePostReq,
) ([]serverapi.APIDeduplicateArchivePostOKItem, error) {
	data, err := c.deduplicateUseCases.ArchiveEntryPercentage(ctx, req.Data)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   err.Error(),
		}
	}

	return pkg.Map(data, func(raw core.DeduplicateArchiveResult) serverapi.APIDeduplicateArchivePostOKItem {
		return serverapi.APIDeduplicateArchivePostOKItem{
			BookID:                 raw.TargetBookID,
			BookOriginURL:          apiservercore.OptURL(raw.OriginBookURL),
			EntryPercentage:        raw.EntryPercentage,
			ReverseEntryPercentage: raw.ReverseEntryPercentage,
		}
	}), nil
}
