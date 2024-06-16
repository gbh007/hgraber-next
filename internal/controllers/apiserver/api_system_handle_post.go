package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APISystemHandlePost(ctx context.Context, req *server.APISystemHandlePostReq) (server.APISystemHandlePostRes, error) {
	result, err := c.parseUseCases.NewBooks(ctx, req.Urls)
	if err != nil {
		return &server.APISystemHandlePostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APISystemHandlePostOK{
		TotalCount:     result.TotalCount,
		LoadedCount:    result.LoadedCount,
		DuplicateCount: result.DuplicateCount,
		ErrorCount:     result.ErrorCount,
		NotHandled:     result.NotHandled,
		Details: pkg.Map(result.Details, func(b entities.BookHandleResult) server.APISystemHandlePostOKDetailsItem {
			return server.APISystemHandlePostOKDetailsItem{
				URL:         b.URL,
				IsDuplicate: b.IsDuplicate,
				// DuplicateID: , // FIXME: заполнять
				IsHandled:   b.IsHandled,
				ErrorReason: server.NewOptString(b.ErrorReason),
			}
		}),
	}, nil
}
