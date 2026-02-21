package fshandlers

import (
	"context"
	"errors"
	"mime"
	"net/http"
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
) (*serverapi.APIFileIDGetOKHeaders, error) {
	fileID, err := uuid.Parse(strings.TrimSuffix(params.ID, path.Ext(params.ID)))
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusBadRequest,
			InnerCode: apiservercore.ValidationCode,
			Details:   err.Error(),
		}
	}

	var fsID *uuid.UUID

	if params.Fsid.Set {
		fsID = &params.Fsid.Value
	}

	body, err := c.fsUseCases.File(ctx, fileID, fsID)
	if errors.Is(err, core.ErrFileNotFound) {
		return nil, apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   err.Error(),
		}
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
