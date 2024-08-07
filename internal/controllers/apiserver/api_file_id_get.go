package apiserver

import (
	"context"
	"errors"
	"mime"
	"net/url"
	"path"
	"strings"

	"github.com/google/uuid"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

func (c *Controller) APIFileIDGet(ctx context.Context, params server.APIFileIDGetParams) (server.APIFileIDGetRes, error) {
	fileID, err := uuid.Parse(strings.TrimSuffix(params.ID, path.Ext(params.ID)))
	if err != nil {
		return &server.APIFileIDGetBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	body, err := c.webAPIUseCases.File(ctx, fileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &server.APIFileIDGetNotFound{
			InnerCode: WebAPIUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIFileIDGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	// Это не самый правильный и ленивый костыль, но пока его будет достаточно
	contentType := mime.TypeByExtension(path.Ext(params.ID))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &server.APIFileIDGetOKHeaders{
		ContentType: contentType,
		Response: server.APIFileIDGetOK{
			Data: body,
		},
	}, nil
}

func (c *Controller) getFileURL(fileID uuid.UUID, ext string) url.URL {
	return url.URL{
		Scheme: c.externalServerScheme,
		Host:   c.externalServerHostWithPort,
		Path:   "/api/file/" + fileID.String() + ext,
	}
}
