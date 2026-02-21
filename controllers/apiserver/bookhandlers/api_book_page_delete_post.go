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
	var (
		err error
		uc  string
	)

	switch req.Type {
	case serverapi.APIBookPageDeletePostReqTypeOne:
		err = c.bookUseCases.DeletePage(ctx, req.BookID, req.PageNumber)
		uc = apiservercore.BookUseCaseCode

	case serverapi.APIBookPageDeletePostReqTypeAllCopy:
		err = c.deduplicateUseCases.DeleteAllPageByHash(ctx, req.BookID, req.PageNumber, req.SetDeadHash.Value)
		uc = apiservercore.DeduplicateUseCaseCode

	default:
		err = fmt.Errorf("unsupported type: %v", req.Type) //nolint:revive // правило не применимо
		uc = apiservercore.ValidationCode
	}

	if errors.Is(err, core.ErrBookNotFound) ||
		errors.Is(err, core.ErrPageNotFound) {
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
