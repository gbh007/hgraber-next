package bookhandlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *BookHandlersController) APIBookDeletePost(
	ctx context.Context,
	req *serverapi.APIBookDeletePostReq,
) error {
	var err error

	switch req.Type {
	case serverapi.APIBookDeletePostReqTypeSoft:
		err = c.bookUseCases.DeleteBook(ctx, req.BookID)

	case serverapi.APIBookDeletePostReqTypePageAndCopy:
		err = c.deduplicateUseCases.RemoveBookPagesWithDeadHash(ctx, req.BookID, req.MarkAsDeletedEmptyBook.Value)

	case serverapi.APIBookDeletePostReqTypeDeadHashedPages:
		err = c.deduplicateUseCases.DeleteBookDeadHashedPages(ctx, req.BookID)

	default:
		err = fmt.Errorf("unsupported type: %v", req.Type)

		return apiservercore.APIError{
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		}
	}

	if errors.Is(err, core.ErrBookNotFound) {
		return apiservercore.APIError{
			Code:    http.StatusNotFound,
			Details: err.Error(),
		}
	}

	if err != nil {
		return err
	}

	return nil
}
