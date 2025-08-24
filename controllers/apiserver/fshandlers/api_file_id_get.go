package fshandlers

import (
	"context"
	"errors"
	"mime"
	"path"
	"strings"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIFileIDGet(
	ctx context.Context,
	params serverapi.APIFileIDGetParams,
) (serverapi.APIFileIDGetRes, error) {
	fileID, err := uuid.Parse(strings.TrimSuffix(params.ID, path.Ext(params.ID)))
	if err != nil {
		return &serverapi.APIFileIDGetBadRequest{
			InnerCode: apiservercore.ValidationCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	var fsID *uuid.UUID

	if params.Fsid.Set {
		fsID = &params.Fsid.Value
	}

	body, err := c.fsUseCases.File(ctx, fileID, fsID)
	if errors.Is(err, core.FileNotFoundError) {
		return &serverapi.APIFileIDGetNotFound{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIFileIDGetInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	// Это не самый правильный и ленивый костыль, но пока его будет достаточно
	contentType := mime.TypeByExtension(path.Ext(params.ID))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &serverapi.APIFileIDGetOKHeaders{
		ContentType: contentType,
		Response: serverapi.APIFileIDGetOK{
			Data: body,
		},
	}, nil
}
