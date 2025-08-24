package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeColorCreatePost(
	ctx context.Context,
	req *serverapi.APIAttributeColorCreatePostReq,
) (serverapi.APIAttributeColorCreatePostRes, error) {
	err := c.attributeUseCases.CreateAttributeColor(ctx, core.AttributeColor{
		Code:            req.Code,
		Value:           req.Value,
		TextColor:       req.TextColor,
		BackgroundColor: req.BackgroundColor,
	})
	if err != nil {
		return &serverapi.APIAttributeColorCreatePostInternalServerError{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeColorCreatePostNoContent{}, nil
}
