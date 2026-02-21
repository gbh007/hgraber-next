package bookhandlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookRawPost(
	ctx context.Context,
	req *serverapi.APIBookRawPostReq,
) (*serverapi.BookRaw, error) {
	var (
		book      core.BookContainer
		err       error
		innerCode string
	)

	switch {
	case req.ID.IsSet():
		innerCode = apiservercore.WebAPIUseCaseCode
		book, err = c.bookUseCases.BookRaw(ctx, req.ID.Value)

	case req.URL.IsSet():
		innerCode = apiservercore.ParseUseCaseCode
		book, err = c.parseUseCases.BookByURL(ctx, req.URL.Value)

	default:
		return nil, apiservercore.APIError{
			Code:      http.StatusBadRequest,
			InnerCode: apiservercore.ValidationCode,
			Details:   "id and url is empty",
		}
	}

	if errors.Is(err, core.ErrBookNotFound) {
		return nil, apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: innerCode,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: innerCode,
			Details:   err.Error(),
		}
	}

	return apiservercore.ConvertBookFullToBookRaw(book), nil
}
