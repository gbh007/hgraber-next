package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemRPCRemoveMismatchFilesPost(ctx context.Context) (serverAPI.APISystemRPCRemoveMismatchFilesPostRes, error) {
	removeFromFS, removeFromDB, err := c.cleanupUseCases.RemoveFilesInStoragesMismatch(ctx)
	if err != nil {
		return &serverAPI.APISystemRPCRemoveMismatchFilesPostInternalServerError{
			InnerCode: CleanupUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemRPCRemoveMismatchFilesPostOK{
		RemoveFromDb: removeFromDB,
		RemoveFromFs: removeFromFS,
	}, nil
}
