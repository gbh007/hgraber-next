package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APISystemTaskCreatePost(ctx context.Context, req *serverapi.APISystemTaskCreatePostReq) (serverapi.APISystemTaskCreatePostRes, error) {
	var code systemmodel.TaskCode

	switch req.Code {
	case "deduplicate_files":
		code = systemmodel.DeduplicateFilesTaskCode
	case "remove_detached_files":
		code = systemmodel.RemoveDetachedFilesTaskCode
	// FIXME: удалить если не будет дальнейших модификаций
	// case "remove_mismatch_files":
	// 	code = entities.RemoveFilesInStoragesMismatchTaskCode
	case "fill_dead_hashes":
		code = systemmodel.FillDeadHashesTaskCode
	case "fill_dead_hashes_with_remove_deleted_pages":
		code = systemmodel.FillDeadHashesAndRemoveDeletedPagesTaskCode
	case "clean_deleted_pages":
		code = systemmodel.CleanDeletedPagesTaskCode
	case "clean_deleted_rebuilds":
		code = systemmodel.CleanDeletedRebuildsTaskCode
	}

	err := c.taskUseCases.RunTask(ctx, code)
	if err != nil {
		return &serverapi.APISystemTaskCreatePostInternalServerError{
			InnerCode: apiservercore.TaskerUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APISystemTaskCreatePostNoContent{}, nil
}
