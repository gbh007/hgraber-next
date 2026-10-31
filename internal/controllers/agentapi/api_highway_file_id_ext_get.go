package agentapi

import (
	"context"
	"errors"
	"mime"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIHighwayFileIDExtGet(ctx context.Context, params agentapi.APIHighwayFileIDExtGetParams) (*agentapi.APIHighwayFileIDExtGetOKHeaders, error) {
	if c.highwayUseCase == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	if params.Token == "" {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "missing token",
		}
	}

	err := c.highwayUseCase.ValidateToken(ctx, params.Token)
	if err != nil {
		return nil, apiError{
			Code:    http.StatusForbidden,
			Details: err.Error(),
		}
	}

	body, err := c.highwayUseCase.Get(ctx, params.ID)
	if errors.Is(err, core.ErrFileNotFound) {
		return nil, apiError{
			Code:    http.StatusNotFound,
			Details: err.Error(),
		}
	}

	if err != nil {
		return nil, err
	}

	// Это не самый правильный и ленивый костыль, но пока его будет достаточно
	contentType := mime.TypeByExtension("." + params.Ext)

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &agentapi.APIHighwayFileIDExtGetOKHeaders{
		ContentType: contentType,
		Response: agentapi.APIHighwayFileIDExtGetOK{
			Data: body,
		},
	}, nil
}
