package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIAttributeCountGet(ctx context.Context) (serverAPI.APIAttributeCountGetRes, error) {
	attributes, err := c.webAPIUseCases.AttributesCount(ctx)
	if err != nil {
		return &serverAPI.APIAttributeCountGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAttributeCountGetOK{
		Attributes: pkg.Map(attributes, func(raw entities.AttributeVariant) serverAPI.APIAttributeCountGetOKAttributesItem {
			return serverAPI.APIAttributeCountGetOKAttributesItem{
				Code:  raw.Code,
				Value: raw.Value,
				Count: raw.Count,
			}
		}),
	}, nil
}
