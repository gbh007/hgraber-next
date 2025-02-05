package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsCreatePost(ctx context.Context, req agentapi.APIFsCreatePostReq, params agentapi.APIFsCreatePostParams) (agentapi.APIFsCreatePostRes, error) {
	return &agentapi.APIFsCreatePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentapi.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsDeletePost(ctx context.Context, req *agentapi.APIFsDeletePostReq) (agentapi.APIFsDeletePostRes, error) {
	return &agentapi.APIFsDeletePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentapi.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsGetGet(ctx context.Context, params agentapi.APIFsGetGetParams) (agentapi.APIFsGetGetRes, error) {
	return &agentapi.APIFsGetGetBadRequest{
		InnerCode: ValidationCode,
		Details:   agentapi.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIParsingBookMultiPost(ctx context.Context, req *agentapi.APIParsingBookMultiPostReq) (agentapi.APIParsingBookMultiPostRes, error) {
	return &agentapi.APIParsingBookMultiPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentapi.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIFsInfoPost(ctx context.Context, req *agentapi.APIFsInfoPostReq) (agentapi.APIFsInfoPostRes, error) {
	return &agentapi.APIFsInfoPostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentapi.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIHighwayFileIDExtGet(ctx context.Context, params agentapi.APIHighwayFileIDExtGetParams) (agentapi.APIHighwayFileIDExtGetRes, error) {
	return &agentapi.APIHighwayFileIDExtGetBadRequest{
		InnerCode: ValidationCode,
		Details:   agentapi.NewOptString("unsupported api"),
	}, nil
}

func (c *Controller) APIHighwayTokenCreatePost(ctx context.Context) (agentapi.APIHighwayTokenCreatePostRes, error) {
	return &agentapi.APIHighwayTokenCreatePostBadRequest{
		InnerCode: ValidationCode,
		Details:   agentapi.NewOptString("unsupported api"),
	}, nil
}
