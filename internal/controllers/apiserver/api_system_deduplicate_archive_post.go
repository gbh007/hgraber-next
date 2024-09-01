package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemDeduplicateArchivePost(ctx context.Context, req serverAPI.APISystemDeduplicateArchivePostReq) (serverAPI.APISystemDeduplicateArchivePostRes, error) {
	data, err := c.deduplicateUseCases.ArchiveEntryPercentage(ctx, req.Data)
	if err != nil {
		return &serverAPI.APISystemDeduplicateArchivePostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	result := serverAPI.APISystemDeduplicateArchivePostOKApplicationJSON(pkg.Map(data, func(raw entities.DeduplicateArchiveResult) serverAPI.APISystemDeduplicateArchivePostOKItem {
		return serverAPI.APISystemDeduplicateArchivePostOKItem{
			BookID:                 raw.TargetBookID,
			BookOriginURL:          optURL(raw.OriginBookURL),
			EntryPercentage:        raw.EntryPercentage,
			ReverseEntryPercentage: raw.ReverseEntryPercentage,
		}
	}))

	return &result, nil
}
