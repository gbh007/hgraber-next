package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorCreatePost(
	ctx context.Context,
	req *serverapi.APIAttributeColorCreatePostReq,
) error {
	return c.attributeUseCases.CreateAttributeColor(ctx, core.AttributeColor{
		Code:            req.Code,
		Value:           req.Value,
		TextColor:       req.TextColor,
		BackgroundColor: req.BackgroundColor,
	})
}
