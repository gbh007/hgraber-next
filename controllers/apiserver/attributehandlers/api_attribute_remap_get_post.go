package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapGetPost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapGetPostReq,
) (*serverapi.AttributeRemap, error) {
	ar, err := c.attributeUseCases.AttributeRemap(ctx, req.Code, req.Value)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	result := apiservercore.ConvertAttributeRemapToAPI(ar)

	return &result, nil
}
