package masterapi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func enrichError(err error) error {
	var errResp *serverapi.ErrorResponseStatusCode

	if errors.As(err, &errResp) {
		switch errResp.StatusCode {
		case http.StatusBadRequest:
			return fmt.Errorf(
				"%w: %s",
				entities.MasterAPIBadRequest,
				errResp.Response.Details.Value,
			)
		case http.StatusUnauthorized:
			return fmt.Errorf(
				"%w: %s",
				entities.MasterAPIUnauthorized,
				errResp.Response.Details.Value,
			)

		case http.StatusForbidden:
			return fmt.Errorf(
				"%w: %s",
				entities.MasterAPIForbidden,
				errResp.Response.Details.Value,
			)

		case http.StatusConflict:
			return fmt.Errorf(
				"%w: %s",
				entities.MasterAPIConflict,
				errResp.Response.Details.Value,
			)

		case http.StatusInternalServerError:
			return fmt.Errorf(
				"%w: %s",
				entities.MasterAPIInternalError,
				errResp.Response.Details.Value,
			)
		}

		return fmt.Errorf(
			"%w: %s",
			entities.MasterAPIUnknownResponse,
			errResp.Response.Details.Value,
		)
	}

	return fmt.Errorf("master api: %w", err)
}
