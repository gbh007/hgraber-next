package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AttributeHandlersController) APIAttributeCountGet(
	ctx context.Context,
) (*serverapi.APIAttributeCountGetOK, error) {
	attributes, err := c.attributeUseCases.AttributesCount(ctx)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
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
