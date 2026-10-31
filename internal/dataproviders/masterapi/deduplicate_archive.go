package masterapi

import (
	"context"
	"io"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *Client) DeduplicateArchive(ctx context.Context, body io.Reader) ([]entities.DeduplicateArchiveResult, error) {
	res, err := c.rawClient.APIDeduplicateArchivePost(ctx, serverapi.APIDeduplicateArchivePostReq{
		Data: body,
	})
	if err != nil {
		return nil, enrichError(err)
	}

	return pkg.Map(res, func(raw serverapi.APIDeduplicateArchivePostOKItem) entities.DeduplicateArchiveResult {
		var u *url.URL

		if raw.BookOriginURL.IsSet() {
			u = &raw.BookOriginURL.Value
		}

		return entities.DeduplicateArchiveResult{
			TargetBookID:           raw.BookID,
			OriginBookURL:          u,
			EntryPercentage:        raw.EntryPercentage,
			ReverseEntryPercentage: raw.ReverseEntryPercentage,
		}
	}), nil
}
