package fshandlers

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIBookPageBodyPost(
	ctx context.Context,
	req *serverapi.APIBookPageBodyPostReq,
) (serverapi.APIBookPageBodyPostOK, error) {
	var (
		body io.Reader
		err  error
	)

	switch {
	case req.ID.IsSet() && req.PageNumber.IsSet():
		body, err = c.fsUseCases.PageBody(ctx, req.ID.Value, req.PageNumber.Value)

	case req.URL.IsSet():
		body, err = c.parseUseCases.PageBodyByURL(ctx, req.URL.Value)

	default:
		return serverapi.APIBookPageBodyPostOK{}, apiservercore.APIError{
			Code:    http.StatusBadRequest,
			Details: "id/page number and url is empty",
		}
	}

	if errors.Is(err, core.ErrPageNotFound) || errors.Is(err, core.ErrFileNotFound) {
		return serverapi.APIBookPageBodyPostOK{}, apiservercore.APIError{
			Code:    http.StatusNotFound,
			Details: err.Error(),
		}
	}

	if err != nil {
		return serverapi.APIBookPageBodyPostOK{}, err //nolint:wrapcheck // будет исправлено позднее
	}

	return serverapi.APIBookPageBodyPostOK{
		Data: body,
	}, nil
}
