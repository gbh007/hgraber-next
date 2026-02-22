package attributehandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AttributeHandlersController) APIAttributeRemapCreatePost(
	ctx context.Context,
	req *serverapi.APIAttributeRemapCreatePostReq,
) error {
	if !(req.IsDelete.Value || (req.ToCode.IsSet() && req.ToValue.IsSet())) { //nolint:staticcheck,lll // будет исправлено позднее
		return apiservercore.APIError{
			Code:    http.StatusBadRequest,
			Details: "invalid remap",
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

	return c.attributeUseCases.CreateAttributeRemap(ctx, ar)
}
