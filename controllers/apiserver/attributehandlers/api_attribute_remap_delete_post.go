package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapDeletePost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapDeletePostReq,
) (serverapi.APIAttributeRemapDeletePostRes, error) {
	err := c.attributeUseCases.DeleteAttributeRemap(ctx, req.Code, req.Value)
	if err != nil {
		return &serverapi.APIAttributeRemapDeletePostInternalServerError{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeRemapDeletePostNoContent{}, nil
}
