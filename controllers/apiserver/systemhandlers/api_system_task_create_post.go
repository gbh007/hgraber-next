package systemhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APISystemTaskCreatePost(
	ctx context.Context,
	req *serverapi.APISystemTaskCreatePostReq,
) error {
	var code systemmodel.TaskCode

	switch req.Code {
	case "deduplicate_files":
		code = systemmodel.DeduplicateFilesTaskCode
	case "remove_detached_files":
		code = systemmodel.RemoveDetachedFilesTaskCode
	case "fill_dead_hashes":
		code = systemmodel.FillDeadHashesTaskCode
	case "fill_dead_hashes_with_remove_deleted_pages":
		code = systemmodel.FillDeadHashesAndRemoveDeletedPagesTaskCode
	case "clean_deleted_pages":
		code = systemmodel.CleanDeletedPagesTaskCode
	case "clean_deleted_rebuilds":
		code = systemmodel.CleanDeletedRebuildsTaskCode
	case "remap_attributes":
		code = systemmodel.RemapAttributesTaskCode
	case "clean_after_rebuild":
		code = systemmodel.CleanAfterRebuildTaskCode
	case "clean_after_parse":
		code = systemmodel.CleanAfterParseTaskCode
	}

	err := c.systemUseCases.RunTask(ctx, code)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.TaskerUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
