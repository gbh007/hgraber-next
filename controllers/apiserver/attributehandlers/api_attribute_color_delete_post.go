package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorDeletePost(
	ctx context.Context,
	req *serverapi.APIAttributeColorDeletePostReq,
) error {
	err := c.attributeUseCases.DeleteAttributeColor(ctx, req.Code, req.Value)
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
