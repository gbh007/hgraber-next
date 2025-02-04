package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIBookRawPost(ctx context.Context, req *serverAPI.APIBookRawPostReq) (serverAPI.APIBookRawPostRes, error) {
	var (
		book      core.BookContainer
		err       error
		innerCode string
	)

	switch {
	case req.ID.IsSet():
		innerCode = WebAPIUseCaseCode
		book, err = c.webAPIUseCases.BookRaw(ctx, req.ID.Value)

	case req.URL.IsSet():
		innerCode = ParseUseCaseCode
		book, err = c.parseUseCases.BookByURL(ctx, req.URL.Value)

	default:
		return &serverAPI.APIBookRawPostBadRequest{
			InnerCode: ValidationCode,
			Details:   serverAPI.NewOptString("id and url is empty"),
		}, nil
	}

	if errors.Is(err, core.BookNotFoundError) {
		return &serverAPI.APIBookRawPostNotFound{
			InnerCode: innerCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIBookRawPostInternalServerError{
			InnerCode: innerCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return convertBookFullToBookRaw(book), nil
}
