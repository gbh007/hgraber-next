package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIParsingPageExistsPost(ctx context.Context, req *serverAPI.APIParsingPageExistsPostReq) (serverAPI.APIParsingPageExistsPostRes, error) {
	result, err := c.parseUseCases.PagesExists(ctx, pkg.Map(req.Urls, func(u serverAPI.APIParsingPageExistsPostReqUrlsItem) agentmodel.AgentPageURL {
		return agentmodel.AgentPageURL{
			BookURL:  u.BookURL,
			ImageURL: u.ImageURL,
		}
	}))
	if err != nil {
		return &serverAPI.APIParsingPageExistsPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIParsingPageExistsPostOK{
		Result: pkg.Map(result, func(v agentmodel.AgentPageCheckResult) serverAPI.APIParsingPageExistsPostOKResultItem {
			switch {
			case v.IsPossible:
				return serverAPI.APIParsingPageExistsPostOKResultItem{
					BookURL:  v.BookURL,
					ImageURL: v.ImageURL,
					Result:   serverAPI.APIParsingPageExistsPostOKResultItemResultOk,
				}

			case v.IsUnsupported:
				return serverAPI.APIParsingPageExistsPostOKResultItem{
					BookURL:  v.BookURL,
					ImageURL: v.ImageURL,
					Result:   serverAPI.APIParsingPageExistsPostOKResultItemResultUnsupported,
				}

			case v.HasError:
				return serverAPI.APIParsingPageExistsPostOKResultItem{
					BookURL:      v.BookURL,
					ImageURL:     v.ImageURL,
					Result:       serverAPI.APIParsingPageExistsPostOKResultItemResultError,
					ErrorDetails: serverAPI.NewOptString(v.ErrorReason),
				}

			default:
				return serverAPI.APIParsingPageExistsPostOKResultItem{
					BookURL:      v.BookURL,
					ImageURL:     v.ImageURL,
					Result:       serverAPI.APIParsingPageExistsPostOKResultItemResultError,
					ErrorDetails: serverAPI.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
