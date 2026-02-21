package bookhandlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookArchiveIDGet(
	ctx context.Context,
	params serverapi.APIBookArchiveIDGetParams,
) (*serverapi.APIBookArchiveIDGetOKHeaders, error) {
	body, book, err := c.exportUseCases.ExportBook(ctx, params.ID)
	if errors.Is(err, core.ErrBookNotFound) {
		return nil, apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APIBookArchiveIDGetOKHeaders{
		ContentDisposition: serverapi.NewOptString("attachment; filename=\"" + book.Filename() + "\""),
		ContentType:        "application/zip",
		Response: serverapi.APIBookArchiveIDGetOK{
			Data: body,
		},
	}, nil
}
