package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapDeletePost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapDeletePostReq,
) error {
	return c.attributeUseCases.DeleteAttributeRemap(ctx, req.Code, req.Value)
}
