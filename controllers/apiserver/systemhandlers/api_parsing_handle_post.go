package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APIParsingHandlePost(
	ctx context.Context,
	req *serverapi.APIParsingHandlePostReq,
) (serverapi.APIParsingHandlePostRes, error) {
	if req.IsMulti.Value {
		result, err := c.parseUseCases.NewBooksMulti(ctx, req.Urls, parsing.ParseFlags{
			AutoVerify: req.AutoVerify.Value,
			ReadOnly:   req.ReadOnlyMode.Value,
		})
		if err != nil {
			return &serverapi.APIParsingHandlePostInternalServerError{
				InnerCode: apiservercore.ParseUseCaseCode,
				Details:   serverapi.NewOptString(err.Error()),
			}, nil
		}

		return &serverapi.APIParsingHandlePostOK{
			TotalCount:     result.Details.TotalCount,
			LoadedCount:    result.Details.LoadedCount,
			DuplicateCount: result.Details.DuplicateCount,
			ErrorCount:     result.Details.ErrorCount,
			// Поскольку в запросе адреса для массовой обработки, то как не обработанные отдаем их же.
			NotHandled: result.NotHandled,
			Details:    convertAPIParsingHandlePostOKDetails(result.Details.Details),
		}, nil
	}

	result, err := c.parseUseCases.NewBooks(ctx, req.Urls, parsing.ParseFlags{
		AutoVerify: req.AutoVerify.Value,
		ReadOnly:   req.ReadOnlyMode.Value,
	})
	if err != nil {
		return &serverapi.APIParsingHandlePostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIParsingHandlePostOK{
		TotalCount:     result.TotalCount,
		LoadedCount:    result.LoadedCount,
		DuplicateCount: result.DuplicateCount,
		ErrorCount:     result.ErrorCount,
		NotHandled:     result.NotHandled,
		Details:        convertAPIParsingHandlePostOKDetails(result.Details),
	}, nil
}

func convertAPIParsingHandlePostOKDetails(
	raw []parsing.BookHandleResult,
) []serverapi.APIParsingHandlePostOKDetailsItem {
	return pkg.Map(raw, func(b parsing.BookHandleResult) serverapi.APIParsingHandlePostOKDetailsItem {
		return serverapi.APIParsingHandlePostOKDetailsItem{
			URL:          b.URL,
			IsDuplicate:  b.IsDuplicate,
			DuplicateIds: b.DuplicateIDs,
			IsHandled:    b.IsHandled,
			ID:           apiservercore.OptUUID(b.ID),
			ErrorReason:  serverapi.NewOptString(b.ErrorReason),
		}
	})
}
