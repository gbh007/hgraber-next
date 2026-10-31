package agentapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIImportArchivePost(ctx context.Context, req agentapi.APIImportArchivePostReq, params agentapi.APIImportArchivePostParams) error {
	if c.importUseCase == nil {
		return apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	var u *url.URL

	if params.BookURL.IsSet() {
		u = &params.BookURL.Value
	}

	err := c.importUseCase.Create(ctx, agentdto.ImportData{
		BookID:   params.BookID,
		BookName: params.BookName,
		Body:     req.Data,
		BookURL:  u,
	})
	if err != nil {
		return err
	}

	return nil
}
