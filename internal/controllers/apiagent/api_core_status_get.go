package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (c *Controller) APICoreStatusGet(ctx context.Context) (agentAPI.APICoreStatusGetRes, error) {
	return &agentAPI.APICoreStatusGetOK{
		StartAt: c.startAt,
		Status:  agentAPI.APICoreStatusGetOKStatusOk,
		Problems: []agentAPI.APICoreStatusGetOKProblemsItem{
			{
				Type:    agentAPI.APICoreStatusGetOKProblemsItemTypeInfo,
				Details: "cache agent",
			},
		},
	}, nil
}
