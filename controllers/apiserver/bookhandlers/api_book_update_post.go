package bookhandlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookUpdatePost(
	ctx context.Context,
	req *serverapi.BookRaw,
) error {
	err := c.rebuilderUseCases.UpdateBook(ctx, apiservercore.ConvertBookRawToBookFull(req))

	if errors.Is(err, core.ErrBookNotFound) {
		return apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.RebuilderUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
