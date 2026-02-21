package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapDeletePost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapDeletePostReq,
) error {
	err := c.attributeUseCases.DeleteAttributeRemap(ctx, req.Code, req.Value)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
