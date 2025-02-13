package fshandlers

import (
	"context"
	"errors"
	"io"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *FSHandlersController) APIPageBodyPost(ctx context.Context, req *serverapi.APIPageBodyPostReq) (serverapi.APIPageBodyPostRes, error) {
	var (
		body      io.Reader
		err       error
		innerCode string
	)

	switch {
	case req.ID.IsSet() && req.PageNumber.IsSet():
		innerCode = apiservercore.WebAPIUseCaseCode
		body, err = c.fsUseCases.PageBody(ctx, req.ID.Value, req.PageNumber.Value)

	case req.URL.IsSet():
		innerCode = apiservercore.ParseUseCaseCode
		body, err = c.parseUseCases.PageBodyByURL(ctx, req.URL.Value)

	default:
		return &serverapi.APIPageBodyPostBadRequest{
			InnerCode: apiservercore.ValidationCode,
			Details:   serverapi.NewOptString("id/page number and url is empty"),
		}, nil
	}

	if errors.Is(err, core.PageNotFoundError) || errors.Is(err, core.FileNotFoundError) {
		return &serverapi.APIPageBodyPostNotFound{
			InnerCode: innerCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIPageBodyPostInternalServerError{
			InnerCode: innerCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIPageBodyPostOK{
		Data: body,
	}, nil
}
