package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookArchiveIDGet(ctx context.Context, params serverapi.APIBookArchiveIDGetParams) (serverapi.APIBookArchiveIDGetRes, error) {
	body, book, err := c.exportUseCases.ExportBook(ctx, params.ID)
	if errors.Is(err, core.BookNotFoundError) {
		return &serverapi.APIBookArchiveIDGetNotFound{
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIBookArchiveIDGetInternalServerError{
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIBookArchiveIDGetOKHeaders{
		ContentDisposition: serverapi.NewOptString("attachment; filename=\"" + book.Filename() + "\""),
		ContentType:        "application/zip",
		Response: serverapi.APIBookArchiveIDGetOK{
			Data: body,
		},
	}, nil
}
