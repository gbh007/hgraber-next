package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorGetPost(
	ctx context.Context,
	req *serverapi.APIAttributeColorGetPostReq,
) (*serverapi.AttributeColor, error) {
	color, err := c.attributeUseCases.AttributeColor(ctx, req.Code, req.Value)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	result := serverapi.AttributeColor(color)

	return &result, nil
}
