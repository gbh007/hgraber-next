package adapter

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/open_api/agentAPI"
)

func (a *Adapter) ExportArchive(ctx context.Context, bookID uuid.UUID, bookName string, body io.Reader) error {
	res, err := a.rawClient.APIExportArchivePost(
		ctx,
		agentAPI.APIExportArchivePostReq{
			Data: body,
		},
		agentAPI.APIExportArchivePostParams{
			BookID:   bookID,
			BookName: bookName,
		},
	)
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIExportArchivePostNoContent:
		return nil

	case *agentAPI.APIExportArchivePostBadRequest:
		return fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIExportArchivePostUnauthorized:
		return fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIExportArchivePostForbidden:
		return fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIExportArchivePostInternalServerError:
		return fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentAPIUnknownResponse
	}
}
