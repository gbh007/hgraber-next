package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorUpdatePost(
	ctx context.Context,
	req *serverapi.APIAttributeColorUpdatePostReq,
) (serverapi.APIAttributeColorUpdatePostRes, error) {
	err := c.attributeUseCases.UpdateAttributeColor(ctx, core.AttributeColor{
		Code:            req.Code,
		Value:           req.Value,
		TextColor:       req.TextColor,
		BackgroundColor: req.BackgroundColor,
	})
	if err != nil {
		return &serverapi.APIAttributeColorUpdatePostInternalServerError{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeColorUpdatePostNoContent{}, nil
}
