package apiagent

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsCreatePost(
	ctx context.Context,
	req agentapi.APIFsCreatePostReq,
	params agentapi.APIFsCreatePostParams,
) error {
	return apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIFsDeletePost(
	ctx context.Context,
	req *agentapi.APIFsDeletePostReq,
) error {
	return apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIFsGetGet(
	ctx context.Context,
	params agentapi.APIFsGetGetParams,
) (agentapi.APIFsGetGetOK, error) {
	return agentapi.APIFsGetGetOK{}, apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIParsingBookMultiPost(
	ctx context.Context,
	req *agentapi.APIParsingBookMultiPostReq,
) (*agentapi.BooksCheckResult, error) {
	return nil, apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIFsInfoPost(
	ctx context.Context,
	req *agentapi.APIFsInfoPostReq,
) (*agentapi.APIFsInfoPostOK, error) {
	return nil, apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIHighwayFileIDExtGet(
	ctx context.Context,
	params agentapi.APIHighwayFileIDExtGetParams,
) (*agentapi.APIHighwayFileIDExtGetOKHeaders, error) {
	return nil, apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIHighwayTokenCreatePost(ctx context.Context) (*agentapi.APIHighwayTokenCreatePostOK, error) {
	return nil, apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIHproxyParseBookPost(
	ctx context.Context,
	req *agentapi.APIHproxyParseBookPostReq,
) (*agentapi.APIHproxyParseBookPostOK, error) {
	return nil, apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}

func (c *Controller) APIHproxyParseListPost(
	ctx context.Context,
	req *agentapi.APIHproxyParseListPostReq,
) (*agentapi.APIHproxyParseListPostOK, error) {
	return nil, apiError{
		Code:      http.StatusBadRequest,
		InnerCode: ValidationCode,
		Details:   "unsupported api",
	}
}
