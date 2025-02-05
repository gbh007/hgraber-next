package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeCountGet(ctx context.Context) (serverapi.APIAttributeCountGetRes, error) {
	attributes, err := c.webAPIUseCases.AttributesCount(ctx)
	if err != nil {
		return &serverapi.APIAttributeCountGetInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeCountGetOK{
		Attributes: pkg.Map(attributes, func(raw core.AttributeVariant) serverapi.APIAttributeCountGetOKAttributesItem {
			return serverapi.APIAttributeCountGetOKAttributesItem{
				Code:  raw.Code,
				Value: raw.Value,
				Count: raw.Count,
			}
		}),
	}, nil
}
