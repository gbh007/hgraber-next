package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeOriginCountGet(
	ctx context.Context,
) (*serverapi.APIAttributeOriginCountGetOK, error) {
	attributes, err := c.attributeUseCases.BookOriginAttributesCount(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APIAttributeOriginCountGetOK{
		Attributes: pkg.Map(
			attributes,
			func(raw core.AttributeVariant) serverapi.APIAttributeOriginCountGetOKAttributesItem {
				return serverapi.APIAttributeOriginCountGetOKAttributesItem{
					Code:  raw.Code,
					Value: raw.Value,
					Count: raw.Count,
				}
			},
		),
	}, nil
}
