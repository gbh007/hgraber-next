package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsGetGet(ctx context.Context, params agentapi.APIFsGetGetParams) (agentapi.APIFsGetGetOK, error) {
	if c.fileUseCase == nil {
		return agentapi.APIFsGetGetOK{}, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	body, err := c.fileUseCase.Get(ctx, params.FileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return agentapi.APIFsGetGetOK{}, apiError{
			Code:    http.StatusNotFound,
			Details: err.Error(),
		}
	}

	if err != nil {
		return agentapi.APIFsGetGetOK{}, err
	}

	// FIXME: работать с типом контента как в основном сервере
	return agentapi.APIFsGetGetOK{
		Data: body,
	}, nil
}
