package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
)

func (a *Adapter) ExportArchive(ctx context.Context, data entities.AgentExportData) error {
	res, err := a.rawClient.APIExportArchivePost(
		ctx,
		agentAPI.APIExportArchivePostReq{
			Data: data.Body,
		},
		agentAPI.APIExportArchivePostParams{
			BookID:   data.BookID,
			BookName: data.BookName,
			BookURL:  optURL(data.BookURL),
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

func optURL(u *url.URL) agentAPI.OptURI {
	if u == nil {
		return agentAPI.OptURI{}
	}

	return agentAPI.NewOptURI(*u)
}
