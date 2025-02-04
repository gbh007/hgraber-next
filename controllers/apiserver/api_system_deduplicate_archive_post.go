package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APISystemDeduplicateArchivePost(ctx context.Context, req serverAPI.APISystemDeduplicateArchivePostReq) (serverAPI.APISystemDeduplicateArchivePostRes, error) {
	data, err := c.deduplicateUseCases.ArchiveEntryPercentage(ctx, req.Data)
	if err != nil {
		return &serverAPI.APISystemDeduplicateArchivePostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	result := serverAPI.APISystemDeduplicateArchivePostOKApplicationJSON(pkg.Map(data, func(raw core.DeduplicateArchiveResult) serverAPI.APISystemDeduplicateArchivePostOKItem {
		return serverAPI.APISystemDeduplicateArchivePostOKItem{
			BookID:                 raw.TargetBookID,
			BookOriginURL:          optURL(raw.OriginBookURL),
			EntryPercentage:        raw.EntryPercentage,
			ReverseEntryPercentage: raw.ReverseEntryPercentage,
		}
	}))

	return &result, nil
}
