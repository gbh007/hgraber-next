package adapter

import (
	"context"
	"io"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) PageLoad(ctx context.Context, url serveragentdto.AgentPageURL) (io.Reader, error) {
	res, err := a.rawClient.APIParsingPagePost(ctx, &agentapi.APIParsingPagePostReq{
		BookURL:  url.BookURL,
		ImageURL: url.ImageURL,
	})
	if err != nil {
		return nil, enrichError(err)
	}

	return res.Data, nil
}
