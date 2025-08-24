package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeOriginCountGet(
	ctx context.Context,
) (serverapi.APIAttributeOriginCountGetRes, error) {
	attributes, err := c.attributeUseCases.BookOriginAttributesCount(ctx)
	if err != nil {
		return &serverapi.APIAttributeOriginCountGetInternalServerError{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
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
