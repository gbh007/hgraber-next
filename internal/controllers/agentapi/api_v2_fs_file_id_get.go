package agentapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2FsFileIDGet(
	ctx context.Context,
	params agentapi.APIV2FsFileIDGetParams,
) (agentapi.APIV2FsFileIDGetOK, error) {
	if c.fileUseCase == nil {
		return agentapi.APIV2FsFileIDGetOK{}, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	body, err := c.fileUseCase.Get(ctx, params.ID)
	if errors.Is(err, core.ErrFileNotFound) {
		return agentapi.APIV2FsFileIDGetOK{}, apiError{
			Code:    http.StatusNotFound,
			Details: err.Error(),
		}
	}

	if err != nil {
		return agentapi.APIV2FsFileIDGetOK{}, err
	}

	// FIXME: работать с типом контента как в основном сервере
	return agentapi.APIV2FsFileIDGetOK{
		Data: body,
	}, nil
}
