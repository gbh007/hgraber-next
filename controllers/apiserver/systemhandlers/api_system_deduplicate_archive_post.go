package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APISystemDeduplicateArchivePost(ctx context.Context, req serverapi.APISystemDeduplicateArchivePostReq) (serverapi.APISystemDeduplicateArchivePostRes, error) {
	data, err := c.deduplicateUseCases.ArchiveEntryPercentage(ctx, req.Data)
	if err != nil {
		return &serverapi.APISystemDeduplicateArchivePostInternalServerError{
			InnerCode: apiservercore.DeduplicateUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	result := serverapi.APISystemDeduplicateArchivePostOKApplicationJSON(pkg.Map(data, func(raw core.DeduplicateArchiveResult) serverapi.APISystemDeduplicateArchivePostOKItem {
		return serverapi.APISystemDeduplicateArchivePostOKItem{
			BookID:                 raw.TargetBookID,
			BookOriginURL:          apiservercore.OptURL(raw.OriginBookURL),
			EntryPercentage:        raw.EntryPercentage,
			ReverseEntryPercentage: raw.ReverseEntryPercentage,
		}
	}))

	return &result, nil
}
