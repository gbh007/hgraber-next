package apiagent

import (
	"context"

	"hgnext/open_api/agentAPI"
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
