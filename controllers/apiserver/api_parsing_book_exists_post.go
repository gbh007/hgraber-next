package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIParsingBookExistsPost(ctx context.Context, req *serverAPI.APIParsingBookExistsPostReq) (serverAPI.APIParsingBookExistsPostRes, error) {
	result, err := c.parseUseCases.BooksExists(ctx, req.Urls)
	if err != nil {
		return &serverAPI.APIParsingBookExistsPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIParsingBookExistsPostOK{
		Result: pkg.Map(result, func(v agentmodel.AgentBookCheckResult) serverAPI.APIParsingBookExistsPostOKResultItem {
			switch {
			case v.IsPossible:
				return serverAPI.APIParsingBookExistsPostOKResultItem{
					URL:                v.URL,
					Result:             serverAPI.APIParsingBookExistsPostOKResultItemResultOk,
					PossibleDuplicates: v.PossibleDuplicates,
				}

			case v.IsUnsupported:
				return serverAPI.APIParsingBookExistsPostOKResultItem{
					URL:    v.URL,
					Result: serverAPI.APIParsingBookExistsPostOKResultItemResultUnsupported,
				}

			case v.HasError:
				return serverAPI.APIParsingBookExistsPostOKResultItem{
					URL:          v.URL,
					Result:       serverAPI.APIParsingBookExistsPostOKResultItemResultError,
					ErrorDetails: serverAPI.NewOptString(v.ErrorReason),
				}

			default:
				return serverAPI.APIParsingBookExistsPostOKResultItem{
					URL:          v.URL,
					Result:       serverAPI.APIParsingBookExistsPostOKResultItemResultError,
					ErrorDetails: serverAPI.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
