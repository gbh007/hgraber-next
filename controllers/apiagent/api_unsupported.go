package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

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

func (c *Controller) APIParsingBookMultiPost(ctx context.Context, req *agentAPI.APIParsingBookMultiPostReq) (agentAPI.APIParsingBookMultiPostRes, error) {
	return &agentAPI.APIParsingBookMultiPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsInfoPost(ctx context.Context, req *agentAPI.APIFsInfoPostReq) (agentAPI.APIFsInfoPostRes, error) {
	return &agentAPI.APIFsInfoPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIHighwayFileIDExtGet(ctx context.Context, params agentAPI.APIHighwayFileIDExtGetParams) (agentAPI.APIHighwayFileIDExtGetRes, error) {
	return &agentAPI.APIHighwayFileIDExtGetBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIHighwayTokenCreatePost(ctx context.Context) (agentAPI.APIHighwayTokenCreatePostRes, error) {
	return &agentAPI.APIHighwayTokenCreatePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentAPI.NewOptString("unsupported api"),
	}, nil
}
