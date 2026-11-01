package masterapi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func enrichError(err error) error {
	if errResp, ok := errors.AsType[*serverapi.ErrorResponseStatusCode](err); ok {
		switch errResp.StatusCode {
		case http.StatusBadRequest:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrMasterAPIBadRequest,
				errResp.Response.Details.Value,
			)
		case http.StatusUnauthorized:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrMasterAPIUnauthorized,
				errResp.Response.Details.Value,
			)

		case http.StatusForbidden:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrMasterAPIForbidden,
				errResp.Response.Details.Value,
			)

		case http.StatusConflict:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrMasterAPIConflict,
				errResp.Response.Details.Value,
			)

		case http.StatusInternalServerError:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrMasterAPIInternalError,
				errResp.Response.Details.Value,
			)
		}

		return fmt.Errorf(
			"%w: %s",
			serveragentdto.ErrMasterAPIUnknownResponse,
			errResp.Response.Details.Value,
		)
	}

	return fmt.Errorf("master api: %w", err)
}
