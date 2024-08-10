package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIParsingPageExistsPost(ctx context.Context, req *serverAPI.APIParsingPageExistsPostReq) (serverAPI.APIParsingPageExistsPostRes, error) {
	result, err := c.parseUseCases.PagesExists(ctx, pkg.Map(req.Urls, func(u serverAPI.APIParsingPageExistsPostReqUrlsItem) entities.AgentPageURL {
		return entities.AgentPageURL{
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
		Result: pkg.Map(result, func(v entities.AgentPageCheckResult) serverAPI.APIParsingPageExistsPostOKResultItem {
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
