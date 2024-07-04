package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

func (c *Controller) APISystemRPCDeduplicateFilesPost(ctx context.Context) (server.APISystemRPCDeduplicateFilesPostRes, error) {
	count, size, err := c.deduplicateUseCases.DeduplicateFiles(ctx)
	if err != nil {
		return &server.APISystemRPCDeduplicateFilesPostInternalServerError{
			InnerCode: DeduplicateUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APISystemRPCDeduplicateFilesPostOK{
		Count:      count,
		Size:       size,
		PrettySize: entities.PrettySize(size),
	}, nil
}
