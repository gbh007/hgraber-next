package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsCreatePost(ctx context.Context, req agentapi.APIFsCreatePostReq, params agentapi.APIFsCreatePostParams) error {
	if c.fileUseCase == nil {
		return apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	err := c.fileUseCase.Create(ctx, params.FileID, req.Data)
	if errors.Is(err, entities.FileAlreadyExistsError) {
		return apiError{
			Code:    http.StatusConflict,
			Details: err.Error(),
		}
	}

	if err != nil {
		return err
	}

	return nil
}
