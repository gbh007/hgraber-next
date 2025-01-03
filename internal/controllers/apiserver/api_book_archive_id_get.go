package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookArchiveIDGet(ctx context.Context, params serverAPI.APIBookArchiveIDGetParams) (serverAPI.APIBookArchiveIDGetRes, error) {
	body, book, err := c.exportUseCases.ExportBook(ctx, params.ID)
	if errors.Is(err, entities.BookNotFoundError) {
		return &serverAPI.APIBookArchiveIDGetNotFound{
			InnerCode: ExportUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookArchiveIDGetInternalServerError{
			InnerCode: ExportUseCaseCode,
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
