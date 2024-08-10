package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemHandlePost(ctx context.Context, req *serverAPI.APISystemHandlePostReq) (serverAPI.APISystemHandlePostRes, error) {
	result, err := c.parseUseCases.NewBooks(ctx, req.Urls)
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
		Details: pkg.Map(result.Details, func(b entities.BookHandleResult) serverAPI.APISystemHandlePostOKDetailsItem {
			return serverAPI.APISystemHandlePostOKDetailsItem{
				URL:         b.URL,
				IsDuplicate: b.IsDuplicate,
				// DuplicateID: , // FIXME: заполнять
				IsHandled:   b.IsHandled,
				ErrorReason: serverAPI.NewOptString(b.ErrorReason),
			}
		}),
	}, nil
}
