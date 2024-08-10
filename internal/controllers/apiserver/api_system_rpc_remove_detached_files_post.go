package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemRPCRemoveDetachedFilesPost(ctx context.Context) (serverAPI.APISystemRPCRemoveDetachedFilesPostRes, error) {
	count, size, err := c.cleanupUseCases.RemoveDetachedFiles(ctx)
	if err != nil {
		return &serverAPI.APISystemRPCRemoveDetachedFilesPostInternalServerError{
			InnerCode: CleanupUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemRPCRemoveDetachedFilesPostOK{
		Count:      count,
		Size:       size,
		PrettySize: entities.PrettySize(size),
	}, nil
}
