package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapUpdatePost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapUpdatePostReq,
) error {
	if !(req.IsDelete.Value || (req.ToCode.IsSet() && req.ToValue.IsSet())) { //nolint:staticcheck,lll // будет исправлено позднее
		return apiservercore.APIError{
			Code:      http.StatusBadRequest,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   "invalid remap",
		}
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
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AttributeUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
