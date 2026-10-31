package agentapi

import (
	"context"
	"strings"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2StatusGet(ctx context.Context) (*agentapi.APIV2StatusGetOK, error) {
	return &agentapi.APIV2StatusGetOK{
		StartAt: c.startAt,
		Status:  agentapi.APIV2StatusGetOKStatusOk,
		Problems: []agentapi.APIV2StatusGetOKProblemsItem{
			{
				Type:    agentapi.APIV2StatusGetOKProblemsItemTypeInfo,
				Details: "parsers: " + strings.Join(c.parserCodes, ", "),
			},
			{
				Type:    agentapi.APIV2StatusGetOKProblemsItemTypeInfo,
				Details: "modules: " + strings.Join(c.enabledModules, ", "),
			},
		},
	}, nil
}
