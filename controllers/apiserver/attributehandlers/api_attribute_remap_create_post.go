package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapCreatePost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapCreatePostReq,
) (serverapi.APIAttributeRemapCreatePostRes, error) {
	if !(req.IsDelete.Value || (req.ToCode.IsSet() && req.ToValue.IsSet())) {
		return &serverapi.APIAttributeRemapCreatePostBadRequest{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString("invalid remap"),
		}, nil
	}

	ar := core.AttributeRemap{
		Code:    req.Code,
		Value:   req.Value,
		ToCode:  req.ToCode.Value,
		ToValue: req.ToValue.Value,
	}

	if req.IsDelete.Value {
		ar.ToCode = ""
		ar.ToValue = ""
	}

	err := c.attributeUseCases.CreateAttributeRemap(ctx, ar)
	if err != nil {
		return &serverapi.APIAttributeRemapCreatePostInternalServerError{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeRemapCreatePostNoContent{}, nil
}
