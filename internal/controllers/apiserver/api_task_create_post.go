package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APITaskCreatePost(ctx context.Context, req *serverAPI.APITaskCreatePostReq) (serverAPI.APITaskCreatePostRes, error) {
	var code entities.TaskCode

	switch req.Code {
	case serverAPI.APITaskCreatePostReqCodeDeduplicateFiles:
		code = entities.DeduplicateFilesTaskCode
	case serverAPI.APITaskCreatePostReqCodeRemoveDetachedFiles:
		code = entities.RemoveDetachedFilesTaskCode
	case serverAPI.APITaskCreatePostReqCodeRemoveMismatchFiles:
		code = entities.RemoveFilesInStoragesMismatchTaskCode
	}

	err := c.taskUseCases.RunTask(ctx, code)
	if err != nil {
		return &serverAPI.APITaskCreatePostInternalServerError{
			InnerCode: TaskerUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APITaskCreatePostNoContent{}, nil
}
