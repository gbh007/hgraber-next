package adapter

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) PagesCheck(
	ctx context.Context,
	urls []agentmodel.AgentPageURL,
) ([]agentmodel.AgentPageCheckResult, error) {
	res, err := a.rawClient.APIParsingPageCheckPost(ctx, &agentapi.APIParsingPageCheckPostReq{
		Urls: pkg.Map(urls, func(u agentmodel.AgentPageURL) agentapi.APIParsingPageCheckPostReqUrlsItem {
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
		func(v agentapi.APIParsingPageCheckPostOKResultItem) agentmodel.AgentPageCheckResult {
			return agentmodel.AgentPageCheckResult{
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
