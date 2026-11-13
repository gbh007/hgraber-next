package masterapi

import (
	"context"
	"io"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *Client) DeduplicateArchive(ctx context.Context, body io.ReadCloser) ([]agentdto.DeduplicateArchiveResult, error) {
	defer body.Close()

	res, err := c.rawClient.APIDeduplicateArchivePost(ctx, serverapi.APIDeduplicateArchivePostReq{
		Data: body,
	})
	if err != nil {
		return nil, enrichError(err)
	}

	return lo.Map(res, func(raw serverapi.APIDeduplicateArchivePostOKItem, _ int) agentdto.DeduplicateArchiveResult {
		var u *url.URL

		if raw.BookOriginURL.IsSet() {
			u = &raw.BookOriginURL.Value
		}

		return agentdto.DeduplicateArchiveResult{
			TargetBookID:           raw.BookID,
			OriginBookURL:          u,
			EntryPercentage:        raw.EntryPercentage,
			ReverseEntryPercentage: raw.ReverseEntryPercentage,
		}
	}), nil
}
