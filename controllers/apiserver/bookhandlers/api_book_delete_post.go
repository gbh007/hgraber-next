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
	var (
		err error
		uc  string
	)

	switch req.Type {
	case serverapi.APIBookDeletePostReqTypeSoft:
		err = c.bookUseCases.DeleteBook(ctx, req.BookID)
		uc = apiservercore.BookUseCaseCode

	case serverapi.APIBookDeletePostReqTypePageAndCopy:
		err = c.deduplicateUseCases.RemoveBookPagesWithDeadHash(ctx, req.BookID, req.MarkAsDeletedEmptyBook.Value)
		uc = apiservercore.DeduplicateUseCaseCode

	case serverapi.APIBookDeletePostReqTypeDeadHashedPages:
		err = c.deduplicateUseCases.DeleteBookDeadHashedPages(ctx, req.BookID)
		uc = apiservercore.DeduplicateUseCaseCode

	default:
		err = fmt.Errorf("unsupported type: %v", req.Type) //nolint:revive // правило не применимо
		uc = apiservercore.ValidationCode
	}

	if errors.Is(err, core.ErrBookNotFound) {
		return apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: uc,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: uc,
			Details:   err.Error(),
		}
	}

	return nil
}
