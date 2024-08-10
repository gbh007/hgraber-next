package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APIParsingBookExistsPost(ctx context.Context, req *server.APIParsingBookExistsPostReq) (server.APIParsingBookExistsPostRes, error) {
	result, err := c.parseUseCases.BooksExists(ctx, req.Urls)
	if err != nil {
		return &server.APIParsingBookExistsPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIParsingBookExistsPostOK{
		Result: pkg.Map(result, func(v entities.AgentBookCheckResult) server.APIParsingBookExistsPostOKResultItem {
			switch {
			case v.IsPossible:
				return server.APIParsingBookExistsPostOKResultItem{
					URL:                v.URL,
					Result:             server.APIParsingBookExistsPostOKResultItemResultOk,
					PossibleDuplicates: v.PossibleDuplicates,
				}

			case v.IsUnsupported:
				return server.APIParsingBookExistsPostOKResultItem{
					URL:    v.URL,
					Result: server.APIParsingBookExistsPostOKResultItemResultUnsupported,
				}

			case v.HasError:
				return server.APIParsingBookExistsPostOKResultItem{
					URL:          v.URL,
					Result:       server.APIParsingBookExistsPostOKResultItemResultError,
					ErrorDetails: server.NewOptString(v.ErrorReason),
				}

			default:
				return server.APIParsingBookExistsPostOKResultItem{
					URL:          v.URL,
					Result:       server.APIParsingBookExistsPostOKResultItemResultError,
					ErrorDetails: server.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
