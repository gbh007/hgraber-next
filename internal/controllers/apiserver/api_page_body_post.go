package apiserver

import (
	"context"
	"errors"
	"io"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

func (c *Controller) APIPageBodyPost(ctx context.Context, req *server.APIPageBodyPostReq) (server.APIPageBodyPostRes, error) {
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
		return &server.APIPageBodyPostBadRequest{
			InnerCode: ValidationCode,
			Details:   server.NewOptString("id/page number and url is empty"),
		}, nil
	}

	if errors.Is(err, entities.PageNotFoundError) || errors.Is(err, entities.FileNotFoundError) {
		return &server.APIPageBodyPostNotFound{
			InnerCode: innerCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIPageBodyPostInternalServerError{
			InnerCode: innerCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIPageBodyPostOK{
		Data: body,
	}, nil
}
