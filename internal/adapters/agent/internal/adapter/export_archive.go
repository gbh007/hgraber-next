package adapter

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"

	"hgnext/internal/adapters/agent/internal/client"
	"hgnext/internal/entities"
)

func (a *Adapter) ExportArchive(ctx context.Context, bookID uuid.UUID, bookName string, body io.Reader) error {
	res, err := a.rawClient.APIExportArchivePost(
		ctx,
		client.APIExportArchivePostReq{
			Data: body,
		},
		client.APIExportArchivePostParams{
			BookID:   bookID,
			BookName: bookName,
		},
	)
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *client.APIExportArchivePostNoContent:
		return nil

	case *client.APIExportArchivePostBadRequest:
		return fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIExportArchivePostUnauthorized:
		return fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIExportArchivePostForbidden:
		return fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIExportArchivePostInternalServerError:
		return fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentAPIUnknownResponse
	}
}
