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

func (c *BookHandlersController) APIBookPageDeletePost(
	ctx context.Context,
	req *serverapi.APIBookPageDeletePostReq,
) error {
	var err error

	switch req.Type {
	case serverapi.APIBookPageDeletePostReqTypeOne:
		err = c.bookUseCases.DeletePage(ctx, req.BookID, req.PageNumber)

	case serverapi.APIBookPageDeletePostReqTypeAllCopy:
		err = c.deduplicateUseCases.DeleteAllPageByHash(ctx, req.BookID, req.PageNumber, req.SetDeadHash.Value)

	default:
		err = fmt.Errorf("unsupported type: %v", req.Type)

		return apiservercore.APIError{
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		}
	}

	if errors.Is(err, core.ErrBookNotFound) ||
		errors.Is(err, core.ErrPageNotFound) {
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
