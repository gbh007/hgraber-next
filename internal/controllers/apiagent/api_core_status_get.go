package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APICoreStatusGet(ctx context.Context) (*agentapi.APICoreStatusGetOK, error) {
	return &agentapi.APICoreStatusGetOK{
		StartAt: c.startAt,
		Status:  agentapi.APICoreStatusGetOKStatusOk,
		Problems: []agentapi.APICoreStatusGetOKProblemsItem{
			{
				Type:    agentapi.APICoreStatusGetOKProblemsItemTypeInfo,
				Details: "cache agent",
			},
		},
	}, nil
}
