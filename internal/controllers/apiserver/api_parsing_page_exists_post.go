package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APIParsingPageExistsPost(ctx context.Context, req *server.APIParsingPageExistsPostReq) (server.APIParsingPageExistsPostRes, error) {
	result, err := c.parseUseCases.PagesExists(ctx, pkg.Map(req.Urls, func(u server.APIParsingPageExistsPostReqUrlsItem) entities.AgentPageURL {
		return entities.AgentPageURL{
			BookURL:  u.BookURL,
			ImageURL: u.ImageURL,
		}
	}))
	if err != nil {
		return &server.APIParsingPageExistsPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIParsingPageExistsPostOK{
		Result: pkg.Map(result, func(v entities.AgentPageCheckResult) server.APIParsingPageExistsPostOKResultItem {
			switch {
			case v.IsPossible:
				return server.APIParsingPageExistsPostOKResultItem{
					BookURL:  v.BookURL,
					ImageURL: v.ImageURL,
					Result:   server.APIParsingPageExistsPostOKResultItemResultOk,
				}

			case v.IsUnsupported:
				return server.APIParsingPageExistsPostOKResultItem{
					BookURL:  v.BookURL,
					ImageURL: v.ImageURL,
					Result:   server.APIParsingPageExistsPostOKResultItemResultUnsupported,
				}

			case v.HasError:
				return server.APIParsingPageExistsPostOKResultItem{
					BookURL:      v.BookURL,
					ImageURL:     v.ImageURL,
					Result:       server.APIParsingPageExistsPostOKResultItemResultError,
					ErrorDetails: server.NewOptString(v.ErrorReason),
				}

			default:
				return server.APIParsingPageExistsPostOKResultItem{
					BookURL:      v.BookURL,
					ImageURL:     v.ImageURL,
					Result:       server.APIParsingPageExistsPostOKResultItemResultError,
					ErrorDetails: server.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
