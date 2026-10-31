package masterapi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func enrichError(err error) error {
	var errResp *serverapi.ErrorResponseStatusCode

	if errors.As(err, &errResp) {
		switch errResp.StatusCode {
		case http.StatusBadRequest:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.MasterAPIBadRequest,
				errResp.Response.Details.Value,
			)
		case http.StatusUnauthorized:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.MasterAPIUnauthorized,
				errResp.Response.Details.Value,
			)

		case http.StatusForbidden:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.MasterAPIForbidden,
				errResp.Response.Details.Value,
			)

		case http.StatusConflict:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.MasterAPIConflict,
				errResp.Response.Details.Value,
			)

		case http.StatusInternalServerError:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.MasterAPIInternalError,
				errResp.Response.Details.Value,
			)
		}

		return fmt.Errorf(
			"%w: %s",
			serveragentdto.MasterAPIUnknownResponse,
			errResp.Response.Details.Value,
		)
	}

	return fmt.Errorf("master api: %w", err)
}
