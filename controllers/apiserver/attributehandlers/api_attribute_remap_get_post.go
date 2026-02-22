package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapGetPost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapGetPostReq,
) (*serverapi.AttributeRemap, error) {
	ar, err := c.attributeUseCases.AttributeRemap(ctx, req.Code, req.Value)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	result := apiservercore.ConvertAttributeRemapToAPI(ar)

	return &result, nil
}
