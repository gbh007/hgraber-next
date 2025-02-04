package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APISystemHandlePost(ctx context.Context, req *serverAPI.APISystemHandlePostReq) (serverAPI.APISystemHandlePostRes, error) {
	if req.IsMulti.Value {
		result, err := c.parseUseCases.NewBooksMulti(ctx, req.Urls, req.AutoVerify.Value)
		if err != nil {
			return &serverAPI.APISystemHandlePostInternalServerError{
				InnerCode: ParseUseCaseCode,
				Details:   serverAPI.NewOptString(err.Error()),
			}, nil
		}

		return &serverAPI.APISystemHandlePostOK{
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
		return &serverAPI.APISystemHandlePostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemHandlePostOK{
		TotalCount:     result.TotalCount,
		LoadedCount:    result.LoadedCount,
		DuplicateCount: result.DuplicateCount,
		ErrorCount:     result.ErrorCount,
		NotHandled:     result.NotHandled,
		Details:        convertAPISystemHandlePostOKDetails(result.Details),
	}, nil
}

func convertAPISystemHandlePostOKDetails(raw []core.BookHandleResult) []serverAPI.APISystemHandlePostOKDetailsItem {
	return pkg.Map(raw, func(b core.BookHandleResult) serverAPI.APISystemHandlePostOKDetailsItem {
		return serverAPI.APISystemHandlePostOKDetailsItem{
			URL:         b.URL,
			IsDuplicate: b.IsDuplicate,
			// DuplicateID: , // FIXME: заполнять
			IsHandled:   b.IsHandled,
			ErrorReason: serverAPI.NewOptString(b.ErrorReason),
		}
	})
}
