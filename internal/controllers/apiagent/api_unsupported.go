package apiagent

import (
	"context"

	"hgnext/open_api/agentAPI"
)

func (c *Controller) APIExportArchivePost(ctx context.Context, req agentAPI.APIExportArchivePostReq, params agentAPI.APIExportArchivePostParams) (agentAPI.APIExportArchivePostRes, error) {
	return &agentAPI.APIExportArchivePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsCreatePost(ctx context.Context, req agentAPI.APIFsCreatePostReq, params agentAPI.APIFsCreatePostParams) (agentAPI.APIFsCreatePostRes, error) {
	return &agentAPI.APIFsCreatePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsDeletePost(ctx context.Context, req *agentAPI.APIFsDeletePostReq) (agentAPI.APIFsDeletePostRes, error) {
	return &agentAPI.APIFsDeletePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsGetGet(ctx context.Context, params agentAPI.APIFsGetGetParams) (agentAPI.APIFsGetGetRes, error) {
	return &agentAPI.APIFsGetGetBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsIdsGet(ctx context.Context) (agentAPI.APIFsIdsGetRes, error) {
	return &agentAPI.APIFsIdsGetBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingBookMultiPost(ctx context.Context, req *agentAPI.APIParsingBookMultiPostReq) (agentAPI.APIParsingBookMultiPostRes, error) {
	return &agentAPI.APIParsingBookMultiPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}
