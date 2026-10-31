package adapter

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func enrichError(err error) error {
	var errResp *agentapi.ErrorResponseStatusCode

	if errors.As(err, &errResp) {
		switch errResp.StatusCode {
		case http.StatusBadRequest:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrAgentAPIBadRequest,
				errResp.Response.Details.Value,
			)
		case http.StatusUnauthorized:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrAgentAPIUnauthorized,
				errResp.Response.Details.Value,
			)

		case http.StatusForbidden:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrAgentAPIForbidden,
				errResp.Response.Details.Value,
			)

		case http.StatusConflict:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrAgentAPIConflict,
				errResp.Response.Details.Value,
			)

		case http.StatusInternalServerError:
			return fmt.Errorf(
				"%w: %s",
				serveragentdto.ErrAgentAPIInternalError,
				errResp.Response.Details.Value,
			)
		}

		return fmt.Errorf(
			"%w: %s",
			serveragentdto.ErrAgentAPIUnknownResponse,
			errResp.Response.Details.Value,
		)
	}

	return fmt.Errorf("request: %w", err)
}
