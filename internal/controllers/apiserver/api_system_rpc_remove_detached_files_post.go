package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

func (c *Controller) APISystemRPCRemoveDetachedFilesPost(ctx context.Context) (server.APISystemRPCRemoveDetachedFilesPostRes, error) {
	count, size, err := c.cleanupUseCases.RemoveDetachedFiles(ctx)
	if err != nil {
		return &server.APISystemRPCRemoveDetachedFilesPostInternalServerError{
			InnerCode: CleanupUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APISystemRPCRemoveDetachedFilesPostOK{
		Count:      count,
		Size:       size,
		PrettySize: entities.PrettySize(size),
	}, nil
}
