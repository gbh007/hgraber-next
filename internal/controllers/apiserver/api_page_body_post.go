package apiserver

import (
	"context"
	"errors"
	"io"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIPageBodyPost(ctx context.Context, req *serverAPI.APIPageBodyPostReq) (serverAPI.APIPageBodyPostRes, error) {
	var (
		body      io.Reader
		err       error
		innerCode string
	)

	switch {
	case req.ID.IsSet() && req.PageNumber.IsSet():
		innerCode = WebAPIUseCaseCode
		body, err = c.webAPIUseCases.PageBody(ctx, req.ID.Value, req.PageNumber.Value)

	case req.URL.IsSet():
		innerCode = ParseUseCaseCode
		body, err = c.parseUseCases.PageBodyByURL(ctx, req.URL.Value)

	default:
		return &serverAPI.APIPageBodyPostBadRequest{
			InnerCode: ValidationCode,
			Details:   serverAPI.NewOptString("id/page number and url is empty"),
		}, nil
	}

	if errors.Is(err, entities.PageNotFoundError) || errors.Is(err, entities.FileNotFoundError) {
		return &serverAPI.APIPageBodyPostNotFound{
			InnerCode: innerCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIPageBodyPostInternalServerError{
			InnerCode: innerCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIPageBodyPostOK{
		Data: body,
	}, nil
}
