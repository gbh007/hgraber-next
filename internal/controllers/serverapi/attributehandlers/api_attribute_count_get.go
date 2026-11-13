package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *AttributeHandlersController) APIAttributeCountGet(
	ctx context.Context,
) (*serverapi.APIAttributeCountGetOK, error) {
	attributes, err := c.attributeUseCases.AttributesCount(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIAttributeCountGetOK{
		Attributes: lo.Map(attributes, func(raw core.AttributeVariant, _ int) serverapi.APIAttributeCountGetOKAttributesItem {
			return serverapi.APIAttributeCountGetOKAttributesItem{
				Code:  raw.Code,
				Value: raw.Value,
				Count: raw.Count,
			}
		}),
	}, nil
}
