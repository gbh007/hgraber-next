package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APIParsingPageExistsPost(ctx context.Context, req *serverapi.APIParsingPageExistsPostReq) (serverapi.APIParsingPageExistsPostRes, error) {
	result, err := c.parseUseCases.PagesExists(ctx, pkg.Map(req.Urls, func(u serverapi.APIParsingPageExistsPostReqUrlsItem) agentmodel.AgentPageURL {
		return agentmodel.AgentPageURL{
			BookURL:  u.BookURL,
			ImageURL: u.ImageURL,
		}
	}))
	if err != nil {
		return &serverapi.APIParsingPageExistsPostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIParsingPageExistsPostOK{
		Result: pkg.Map(result, func(v agentmodel.AgentPageCheckResult) serverapi.APIParsingPageExistsPostOKResultItem {
			switch {
			case v.IsPossible:
				return serverapi.APIParsingPageExistsPostOKResultItem{
					BookURL:  v.BookURL,
					ImageURL: v.ImageURL,
					Result:   serverapi.APIParsingPageExistsPostOKResultItemResultOk,
				}

			case v.IsUnsupported:
				return serverapi.APIParsingPageExistsPostOKResultItem{
					BookURL:  v.BookURL,
					ImageURL: v.ImageURL,
					Result:   serverapi.APIParsingPageExistsPostOKResultItemResultUnsupported,
				}

			case v.HasError:
				return serverapi.APIParsingPageExistsPostOKResultItem{
					BookURL:      v.BookURL,
					ImageURL:     v.ImageURL,
					Result:       serverapi.APIParsingPageExistsPostOKResultItemResultError,
					ErrorDetails: serverapi.NewOptString(v.ErrorReason),
				}

			default:
				return serverapi.APIParsingPageExistsPostOKResultItem{
					BookURL:      v.BookURL,
					ImageURL:     v.ImageURL,
					Result:       serverapi.APIParsingPageExistsPostOKResultItemResultError,
					ErrorDetails: serverapi.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
