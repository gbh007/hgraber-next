package api

import (
	"context"
	"strings"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APICoreStatusGet(ctx context.Context) (*agentapi.APICoreStatusGetOK, error) {
	return &agentapi.APICoreStatusGetOK{
		StartAt: c.startAt,
		Status:  agentapi.APICoreStatusGetOKStatusOk,
		Problems: []agentapi.APICoreStatusGetOKProblemsItem{
			{
				Type:    agentapi.APICoreStatusGetOKProblemsItemTypeInfo,
				Details: "parsers: " + strings.Join(c.parserCodes, ", "),
			},
			{
				Type:    agentapi.APICoreStatusGetOKProblemsItemTypeInfo,
				Details: "modules: " + strings.Join(c.enabledModules, ", "),
			},
		},
	}, nil
}
