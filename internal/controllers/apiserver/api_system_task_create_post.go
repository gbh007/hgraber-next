package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APISystemTaskCreatePost(ctx context.Context, req *serverAPI.APISystemTaskCreatePostReq) (serverAPI.APISystemTaskCreatePostRes, error) {
	var code entities.TaskCode

	switch req.Code {
	case "deduplicate_files":
		code = entities.DeduplicateFilesTaskCode
	case "remove_detached_files":
		code = entities.RemoveDetachedFilesTaskCode
	// FIXME: удалить если не будет дальнейших модификаций
	// case "remove_mismatch_files":
	// 	code = entities.RemoveFilesInStoragesMismatchTaskCode
	case "fill_dead_hashes":
		code = entities.FillDeadHashesTaskCode
	case "fill_dead_hashes_with_remove_deleted_pages":
		code = entities.FillDeadHashesAndRemoveDeletedPagesTaskCode
	case "clean_deleted_pages":
		code = entities.CleanDeletedPagesTaskCode
	case "clean_deleted_rebuilds":
		code = entities.CleanDeletedRebuildsTaskCode
	}

	err := c.taskUseCases.RunTask(ctx, code)
	if err != nil {
		return &serverAPI.APISystemTaskCreatePostInternalServerError{
			InnerCode: TaskerUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APISystemTaskCreatePostNoContent{}, nil
}
