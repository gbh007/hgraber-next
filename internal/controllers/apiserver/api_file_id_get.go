package apiserver

import (
	"context"
	"errors"
	"mime"
	"net/url"
	"path"
	"strings"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIFileIDGet(ctx context.Context, params serverAPI.APIFileIDGetParams) (serverAPI.APIFileIDGetRes, error) {
	fileID, err := uuid.Parse(strings.TrimSuffix(params.ID, path.Ext(params.ID)))
	if err != nil {
		return &serverAPI.APIFileIDGetBadRequest{
			InnerCode: ValidationCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	body, err := c.webAPIUseCases.File(ctx, fileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &serverAPI.APIFileIDGetNotFound{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIFileIDGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	// Это не самый правильный и ленивый костыль, но пока его будет достаточно
	contentType := mime.TypeByExtension(path.Ext(params.ID))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &serverAPI.APIFileIDGetOKHeaders{
		ContentType: contentType,
		Response: serverAPI.APIFileIDGetOK{
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
