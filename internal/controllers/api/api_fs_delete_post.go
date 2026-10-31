package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsDeletePost(ctx context.Context, req *agentapi.APIFsDeletePostReq) error {
	if c.fileUseCase == nil {
		return apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	err := c.fileUseCase.Delete(ctx, req.FileID)
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
