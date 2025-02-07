package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APISystemHandlePost(ctx context.Context, req *serverapi.APISystemHandlePostReq) (serverapi.APISystemHandlePostRes, error) {
	if req.IsMulti.Value {
		result, err := c.parseUseCases.NewBooksMulti(ctx, req.Urls, req.AutoVerify.Value)
		if err != nil {
			return &serverapi.APISystemHandlePostInternalServerError{
				InnerCode: apiservercore.ParseUseCaseCode,
				Details:   serverapi.NewOptString(err.Error()),
			}, nil
		}

		return &serverapi.APISystemHandlePostOK{
			TotalCount:     result.Details.TotalCount,
			LoadedCount:    result.Details.LoadedCount,
			DuplicateCount: result.Details.DuplicateCount,
			ErrorCount:     result.Details.ErrorCount,
			NotHandled:     result.NotHandled, // Поскольку в запросе адреса для массовой обработки, то как не обработанные отдаем их же.
			Details:        convertAPISystemHandlePostOKDetails(result.Details.Details),
		}, nil
	}

	result, err := c.parseUseCases.NewBooks(ctx, req.Urls, req.AutoVerify.Value)
	if err != nil {
		return &serverapi.APISystemHandlePostInternalServerError{
			InnerCode: apiservercore.ParseUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APISystemHandlePostOK{
		TotalCount:     result.TotalCount,
		LoadedCount:    result.LoadedCount,
		DuplicateCount: result.DuplicateCount,
		ErrorCount:     result.ErrorCount,
		NotHandled:     result.NotHandled,
		Details:        convertAPISystemHandlePostOKDetails(result.Details),
	}, nil
}

func convertAPISystemHandlePostOKDetails(raw []parsing.BookHandleResult) []serverapi.APISystemHandlePostOKDetailsItem {
	return pkg.Map(raw, func(b parsing.BookHandleResult) serverapi.APISystemHandlePostOKDetailsItem {
		return serverapi.APISystemHandlePostOKDetailsItem{
			URL:         b.URL,
			IsDuplicate: b.IsDuplicate,
			// DuplicateID: , // FIXME: заполнять
			IsHandled:   b.IsHandled,
			ErrorReason: serverapi.NewOptString(b.ErrorReason),
		}
	})
}
