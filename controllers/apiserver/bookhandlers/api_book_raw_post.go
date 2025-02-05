package bookhandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookRawPost(ctx context.Context, req *serverapi.APIBookRawPostReq) (serverapi.APIBookRawPostRes, error) {
	var (
		book      core.BookContainer
		err       error
		innerCode string
	)

	switch {
	case req.ID.IsSet():
		innerCode = apiservercore.WebAPIUseCaseCode
		book, err = c.webAPIUseCases.BookRaw(ctx, req.ID.Value)

	case req.URL.IsSet():
		innerCode = apiservercore.ParseUseCaseCode
		book, err = c.parseUseCases.BookByURL(ctx, req.URL.Value)

	default:
		return &serverapi.APIBookRawPostBadRequest{
			InnerCode: apiservercore.ValidationCode,
			Details:   serverapi.NewOptString("id and url is empty"),
		}, nil
	}

	if errors.Is(err, core.BookNotFoundError) {
		return &serverapi.APIBookRawPostNotFound{
			InnerCode: innerCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIBookRawPostInternalServerError{
			InnerCode: innerCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return apiservercore.ConvertBookFullToBookRaw(book), nil
}
