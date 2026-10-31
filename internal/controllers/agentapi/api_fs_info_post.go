package agentapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/fsmodel"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsInfoPost(ctx context.Context, req *agentapi.APIFsInfoPostReq) (*agentapi.APIFsInfoPostOK, error) {
	if c.fileUseCase == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	state, err := c.fileUseCase.State(ctx, req.IncludeFileIds.Value, req.IncludeFileSizes.Value)
	if err != nil {
		return nil, err
	}

	return &agentapi.APIFsInfoPostOK{
		FileIds: state.FileIDs,
		TotalFileCount: agentapi.OptInt64{
			Value: state.TotalFileCount,
			Set:   state.TotalFileCount > 0,
		},
		TotalFileSize: agentapi.OptInt64{
			Value: state.TotalFileSize,
			Set:   state.TotalFileSize > 0,
		},
		AvailableSize: agentapi.OptInt64{
			Value: state.AvailableSize,
			Set:   state.AvailableSize > 0,
		},
		Files: pkg.Map(state.Files, func(raw fsmodel.FSStateFile) agentapi.APIFsInfoPostOKFilesItem {
			return agentapi.APIFsInfoPostOKFilesItem{
				ID:        raw.ID,
				Size:      raw.Size,
				CreatedAt: raw.CreatedAt,
			}
		}),
	}, nil
}
