package api

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingPageCheckPost(ctx context.Context, req *agentapi.APIParsingPageCheckPostReq) (*agentapi.APIParsingPageCheckPostOK, error) {
	if c.parsingUseCases == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	result, err := c.parsingUseCases.CheckPages(ctx, pkg.Map(req.Urls, func(u agentapi.APIParsingPageCheckPostReqUrlsItem) entities.AgentPageURL {
		return entities.AgentPageURL{
			BookURL:  u.BookURL,
			ImageURL: u.ImageURL,
		}
	}))
	if err != nil {
		return nil, err
	}

	return &agentapi.APIParsingPageCheckPostOK{
		Result: pkg.Map(result, func(p entities.AgentPageCheckResult) agentapi.APIParsingPageCheckPostOKResultItem {
			item := agentapi.APIParsingPageCheckPostOKResultItem{
				BookURL:  p.BookURL,
				ImageURL: p.ImageURL,
			}

			switch {
			case p.HasError:
				item.Result = agentapi.APIParsingPageCheckPostOKResultItemResultError
				item.ErrorDetails = agentapi.NewOptString(p.ErrorReason)

			case p.IsPossible:
				item.Result = agentapi.APIParsingPageCheckPostOKResultItemResultOk

			case p.IsUnsupported:
				item.Result = agentapi.APIParsingPageCheckPostOKResultItemResultUnsupported
			}

			return item
		}),
	}, nil
}
