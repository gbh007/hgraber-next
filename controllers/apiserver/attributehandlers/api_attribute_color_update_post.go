package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorUpdatePost(
	ctx context.Context,
	req *serverapi.APIAttributeColorUpdatePostReq,
) error {
	return c.attributeUseCases.UpdateAttributeColor(ctx, core.AttributeColor{
		Code:            req.Code,
		Value:           req.Value,
		TextColor:       req.TextColor,
		BackgroundColor: req.BackgroundColor,
	})
}
