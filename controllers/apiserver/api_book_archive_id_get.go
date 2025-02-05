package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIBookArchiveIDGet(ctx context.Context, params serverAPI.APIBookArchiveIDGetParams) (serverAPI.APIBookArchiveIDGetRes, error) {
	body, book, err := c.exportUseCases.ExportBook(ctx, params.ID)
	if errors.Is(err, core.BookNotFoundError) {
		return &serverAPI.APIBookArchiveIDGetNotFound{
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookArchiveIDGetInternalServerError{
			InnerCode: apiservercore.ExportUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookArchiveIDGetOKHeaders{
		ContentDisposition: serverAPI.NewOptString("attachment; filename=\"" + book.Filename() + "\""),
		ContentType:        "application/zip",
		Response: serverAPI.APIBookArchiveIDGetOK{
			Data: body,
		},
	}, nil
}
