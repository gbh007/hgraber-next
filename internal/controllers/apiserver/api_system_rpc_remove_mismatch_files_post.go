package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
)

func (c *Controller) APISystemRPCRemoveMismatchFilesPost(ctx context.Context) (server.APISystemRPCRemoveMismatchFilesPostRes, error) {
	removeFromFS, removeFromDB, err := c.cleanupUseCases.RemoveFilesInStoragesMismatch(ctx)
	if err != nil {
		return &server.APISystemRPCRemoveMismatchFilesPostInternalServerError{
			InnerCode: CleanupUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APISystemRPCRemoveMismatchFilesPostOK{
		RemoveFromDb: removeFromDB,
		RemoveFromFs: removeFromFS,
	}, nil
}
