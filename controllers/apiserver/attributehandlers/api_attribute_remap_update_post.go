package attributehandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapUpdatePost(ctx context.Context, req *serverapi.APIAttributeRemapUpdatePostReq) (serverapi.APIAttributeRemapUpdatePostRes, error) {
	if !(req.IsDelete.Value || (req.ToCode.IsSet() && req.ToValue.IsSet())) {
		return &serverapi.APIAttributeRemapUpdatePostBadRequest{
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

	err := c.attributeUseCases.UpdateAttributeRemap(ctx, ar)
	if err != nil {
		return &serverapi.APIAttributeRemapUpdatePostInternalServerError{
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAttributeRemapUpdatePostNoContent{}, nil
}
