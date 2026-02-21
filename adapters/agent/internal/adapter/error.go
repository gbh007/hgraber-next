package adapter

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func enrichError(err error) error {
	var errResp *agentapi.ErrorResponseStatusCode

	if errors.As(err, &errResp) {
		switch errResp.StatusCode {
		case http.StatusBadRequest:
			return fmt.Errorf(
				"%w: %s",
				agentmodel.ErrAgentAPIBadRequest,
				errResp.Response.Details.Value,
			)
		case http.StatusUnauthorized:
			return fmt.Errorf(
				"%w: %s",
				agentmodel.ErrAgentAPIUnauthorized,
				errResp.Response.Details.Value,
			)

		case http.StatusForbidden:
			return fmt.Errorf(
				"%w: %s",
				agentmodel.ErrAgentAPIForbidden,
				errResp.Response.Details.Value,
			)

		case http.StatusConflict:
			return fmt.Errorf(
				"%w: %s",
				agentmodel.ErrAgentAPIConflict,
				errResp.Response.Details.Value,
			)

		case http.StatusInternalServerError:
			return fmt.Errorf(
				"%w: %s",
				agentmodel.ErrAgentAPIInternalError,
				errResp.Response.Details.Value,
			)
		}

		return fmt.Errorf(
			"%w: %s",
			agentmodel.ErrAgentAPIUnknownResponse,
			errResp.Response.Details.Value,
		)
	}

	return fmt.Errorf("request: %w", err)
}
