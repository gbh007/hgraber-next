package hproxyhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *HProxyHandlersController) APIHproxyFileGet(ctx context.Context, params serverapi.APIHproxyFileGetParams) (serverapi.APIHproxyFileGetRes, error) {
	r, err := c.hProxyUseCases.Image(ctx, params.BookURL, params.ImageURL)
	if err != nil {
		return &serverapi.APIHproxyFileGetInternalServerError{
			InnerCode: apiservercore.HProxyUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIHproxyFileGetOKHeaders{
		ContentType: "application/octet-stream", // TODO: работать с типом
		Response: serverapi.APIHproxyFileGetOK{
			Data: r,
		},
	}, nil
}
