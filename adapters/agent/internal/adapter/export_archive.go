package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) ExportArchive(ctx context.Context, data agentmodel.AgentExportData) error {
	res, err := a.rawClient.APIImportArchivePost(
		ctx,
		agentapi.APIImportArchivePostReq{
			Data: data.Body,
		},
		agentapi.APIImportArchivePostParams{
			BookID:   data.BookID,
			BookName: data.BookName,
			BookURL:  optURL(data.BookURL),
		},
	)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}

	switch typedRes := res.(type) {
	case *agentapi.APIImportArchivePostNoContent:
		return nil

	case *agentapi.APIImportArchivePostBadRequest:
		return fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIImportArchivePostUnauthorized:
		return fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIImportArchivePostForbidden:
		return fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIImportArchivePostInternalServerError:
		return fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIInternalError, typedRes.Details.Value)

	default:
		return agentmodel.ErrAgentAPIUnknownResponse
	}
}

func optURL(u *url.URL) agentapi.OptURI {
	if u == nil {
		return agentapi.OptURI{}
	}

	return agentapi.NewOptURI(*u)
}
