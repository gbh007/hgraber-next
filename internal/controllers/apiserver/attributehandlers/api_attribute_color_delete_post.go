package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorDeletePost(
	ctx context.Context,
	req *serverapi.APIAttributeColorDeletePostReq,
) error {
	return c.attributeUseCases.DeleteAttributeColor(ctx, req.Code, req.Value)
}
