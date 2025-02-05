package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) ExportArchive(ctx context.Context, data agentmodel.AgentExportData) error {
	res, err := a.rawClient.APIExportArchivePost(
		ctx,
		agentapi.APIExportArchivePostReq{
			Data: data.Body,
		},
		agentapi.APIExportArchivePostParams{
			BookID:   data.BookID,
			BookName: data.BookName,
			BookURL:  optURL(data.BookURL),
		},
	)
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *agentapi.APIExportArchivePostNoContent:
		return nil

	case *agentapi.APIExportArchivePostBadRequest:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIExportArchivePostUnauthorized:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIExportArchivePostForbidden:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIExportArchivePostInternalServerError:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return agentmodel.AgentAPIUnknownResponse
	}
}

func optURL(u *url.URL) agentapi.OptURI {
	if u == nil {
		return agentapi.OptURI{}
	}

	return agentapi.NewOptURI(*u)
}
