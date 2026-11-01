package agentapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2FsFileIDDelete(
	ctx context.Context,
	params agentapi.APIV2FsFileIDDeleteParams,
) error {
	if c.fileUseCase == nil {
		return apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	err := c.fileUseCase.Delete(ctx, params.ID)
	if errors.Is(err, core.ErrFileNotFound) {
		return apiError{
			Code:    http.StatusNotFound,
			Details: err.Error(),
		}
	}

	if err != nil {
		return err
	}

	return nil
}
