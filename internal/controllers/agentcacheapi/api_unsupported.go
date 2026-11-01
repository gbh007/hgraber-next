package agentcacheapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2FsFileIDPost(
	ctx context.Context,
	req agentapi.APIV2FsFileIDPostReq,
	params agentapi.APIV2FsFileIDPostParams,
) error {
	return apiError{
		Code:    http.StatusBadRequest,
		Details: "unsupported api",
	}
}

func (c *Controller) APIV2FsFileIDDelete(
	ctx context.Context,
	params agentapi.APIV2FsFileIDDeleteParams,
) error {
	return apiError{
		Code:    http.StatusBadRequest,
		Details: "unsupported api",
	}
}

func (c *Controller) APIV2FsFileIDGet(
	ctx context.Context,
	params agentapi.APIV2FsFileIDGetParams,
) (agentapi.APIV2FsFileIDGetOK, error) {
	return agentapi.APIV2FsFileIDGetOK{}, apiError{
		Code:    http.StatusBadRequest,
		Details: "unsupported api",
	}
}

func (c *Controller) APIV2ParsingListPost(
	ctx context.Context,
	req *agentapi.APIV2ParsingListPostReq,
) (*agentapi.APIV2ParsingListPostOK, error) {
	return nil, apiError{
		Code:    http.StatusBadRequest,
		Details: "unsupported api",
	}
}

func (c *Controller) APIV2FsGet(
	ctx context.Context,
	params agentapi.APIV2FsGetParams,
) (*agentapi.APIV2FsGetOK, error) {
	return nil, apiError{
		Code:    http.StatusBadRequest,
		Details: "unsupported api",
	}
}

func (c *Controller) APIV2HighwayFileIDExtGet(
	ctx context.Context,
	params agentapi.APIV2HighwayFileIDExtGetParams,
) (*agentapi.APIV2HighwayFileIDExtGetOKHeaders, error) {
	return nil, apiError{
		Code:    http.StatusBadRequest,
		Details: "unsupported api",
	}
}

func (c *Controller) APIV2HighwayTokenPost(ctx context.Context) (*agentapi.APIV2HighwayTokenPostOK, error) {
	return nil, apiError{
		Code:    http.StatusBadRequest,
		Details: "unsupported api",
	}
}
