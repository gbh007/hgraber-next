package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APIParsingBookExistsPost(ctx context.Context, req *serverapi.APIParsingBookExistsPostReq) (serverapi.APIParsingBookExistsPostRes, error) {
	result, err := c.parseUseCases.BooksExists(ctx, req.Urls)
	if err != nil {
		return &serverapi.APIParsingBookExistsPostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIParsingBookExistsPostOK{
		Result: pkg.Map(result, func(v agentmodel.AgentBookCheckResult) serverapi.APIParsingBookExistsPostOKResultItem {
			switch {
			case v.IsPossible:
				return serverapi.APIParsingBookExistsPostOKResultItem{
					URL:    v.URL,
					Result: serverapi.APIParsingBookExistsPostOKResultItemResultOk,
					// FIXME: попробовать оживить.
					// PossibleDuplicates: v.PossibleDuplicates,
				}

			case v.IsUnsupported:
				return serverapi.APIParsingBookExistsPostOKResultItem{
					URL:    v.URL,
					Result: serverapi.APIParsingBookExistsPostOKResultItemResultUnsupported,
				}

			case v.HasError:
				return serverapi.APIParsingBookExistsPostOKResultItem{
					URL:          v.URL,
					Result:       serverapi.APIParsingBookExistsPostOKResultItemResultError,
					ErrorDetails: serverapi.NewOptString(v.ErrorReason),
				}

			default:
				return serverapi.APIParsingBookExistsPostOKResultItem{
					URL:          v.URL,
					Result:       serverapi.APIParsingBookExistsPostOKResultItemResultError,
					ErrorDetails: serverapi.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
