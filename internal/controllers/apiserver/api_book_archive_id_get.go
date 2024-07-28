package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

func (c *Controller) APIBookArchiveIDGet(ctx context.Context, params server.APIBookArchiveIDGetParams) (server.APIBookArchiveIDGetRes, error) {
	body, err := c.exportUseCases.ExportBook(ctx, params.ID)
	if errors.Is(err, entities.BookNotFoundError) {
		return &server.APIBookArchiveIDGetNotFound{
			InnerCode: ExportUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIBookArchiveIDGetInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIBookArchiveIDGetOK{
		Data: body,
	}, nil
}
