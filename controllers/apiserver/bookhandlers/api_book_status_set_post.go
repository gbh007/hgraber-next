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

func (c *BookHandlersController) APIBookStatusSetPost(
	ctx context.Context,
	req *serverapi.APIBookStatusSetPostReq,
) error {
	var err error

	switch req.Status {
	case serverapi.APIBookStatusSetPostReqStatusRebuild:
		err = c.bookUseCases.SetBookRebuild(ctx, req.ID, req.Value)
	case serverapi.APIBookStatusSetPostReqStatusVerify:
		err = c.bookUseCases.VerifyBook(ctx, req.ID, req.Value)
	default:
		err = fmt.Errorf("unsupported status: %v", req.Status)

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
