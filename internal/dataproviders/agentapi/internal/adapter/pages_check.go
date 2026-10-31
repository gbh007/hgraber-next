package adapter

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) PagesCheck(
	ctx context.Context,
	urls []serveragentdto.AgentPageURL,
) ([]serveragentdto.AgentPageCheckResult, error) {
	res, err := a.rawClient.APIParsingPageCheckPost(ctx, &agentapi.APIParsingPageCheckPostReq{
		Urls: pkg.Map(urls, func(u serveragentdto.AgentPageURL) agentapi.APIParsingPageCheckPostReqUrlsItem {
			return agentapi.APIParsingPageCheckPostReqUrlsItem{
				BookURL:  u.BookURL,
				ImageURL: u.ImageURL,
			}
		}),
	})
	if err != nil {
		return nil, enrichError(err)
	}

	return pkg.Map(
		res.Result,
		func(v agentapi.APIParsingPageCheckPostOKResultItem) serveragentdto.AgentPageCheckResult {
			return serveragentdto.AgentPageCheckResult{
				BookURL:       v.BookURL,
				ImageURL:      v.ImageURL,
				IsUnsupported: v.Result == agentapi.APIParsingPageCheckPostOKResultItemResultUnsupported,
				IsPossible:    v.Result == agentapi.APIParsingPageCheckPostOKResultItemResultOk,
				HasError:      v.Result == agentapi.APIParsingPageCheckPostOKResultItemResultError,
				ErrorReason:   v.ErrorDetails.Value,
			}
		},
	), nil
}
