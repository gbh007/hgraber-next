package hproxyhandlers

import (
	"context"
	"mime"
	"path"

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

	// Это не самый правильный и ленивый костыль, но пока его будет достаточно
	contentType := mime.TypeByExtension(path.Ext(params.ImageURL.String()))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &serverapi.APIHproxyFileGetOKHeaders{
		ContentType: contentType,
		Response: serverapi.APIHproxyFileGetOK{
			Data: r,
		},
	}, nil
}
