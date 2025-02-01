package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIFsValidatePost(ctx context.Context, req *serverAPI.APIFsValidatePostReq) (serverAPI.APIFsValidatePostRes, error) {
	err := c.fsUseCases.ValidateFS(ctx, req.ID)
	if err != nil {
		return &serverAPI.APIFsValidatePostInternalServerError{
			InnerCode: FSUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIFsValidatePostNoContent{}, nil
}
