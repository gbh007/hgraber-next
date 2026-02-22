package hproxyhandlers

import (
	"context"
	"mime"
	"path"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *HProxyHandlersController) APIHproxyFileGet(
	ctx context.Context,
	params serverapi.APIHproxyFileGetParams,
) (*serverapi.APIHproxyFileGetOKHeaders, error) {
	r, err := c.hProxyUseCases.Image(ctx, params.BookURL, params.ImageURL)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
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
