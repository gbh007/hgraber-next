package apiserver

import (
	"context"
	"errors"
	"mime"
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

	var fsID *uuid.UUID

	if params.Fsid.Set {
		fsID = &params.Fsid.Value
	}

	body, err := c.webAPIUseCases.File(ctx, fileID, fsID)
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
