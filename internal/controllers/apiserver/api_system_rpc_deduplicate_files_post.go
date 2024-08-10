package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemRPCDeduplicateFilesPost(ctx context.Context) (serverAPI.APISystemRPCDeduplicateFilesPostRes, error) {
	count, size, err := c.deduplicateUseCases.DeduplicateFiles(ctx)
	if err != nil {
		return &serverAPI.APISystemRPCDeduplicateFilesPostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemRPCDeduplicateFilesPostOK{
		Count:      count,
		Size:       size,
		PrettySize: entities.PrettySize(size),
	}, nil
}
