package adapter

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *FSAdapter) Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error {
	res, err := a.rawClient.APIFsCreatePost(
		ctx,
		agentapi.APIFsCreatePostReq{Data: body},
		agentapi.APIFsCreatePostParams{FileID: fileID},
	)
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *agentapi.APIFsCreatePostNoContent:
		return nil

	case *agentapi.APIFsCreatePostConflict:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIConflict, typedRes.Details.Value)

	case *agentapi.APIFsCreatePostBadRequest:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIFsCreatePostUnauthorized:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIFsCreatePostForbidden:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIFsCreatePostInternalServerError:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) Delete(ctx context.Context, fileID uuid.UUID) error {
	res, err := a.rawClient.APIFsDeletePost(ctx, &agentapi.APIFsDeletePostReq{FileID: fileID})
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *agentapi.APIFsDeletePostNoContent:
		return nil

	case *agentapi.APIFsDeletePostNotFound:
		return fmt.Errorf("%w: %s", core.FileNotFoundError, typedRes.Details.Value)

	case *agentapi.APIFsDeletePostBadRequest:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIFsDeletePostUnauthorized:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIFsDeletePostForbidden:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIFsDeletePostInternalServerError:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	res, err := a.rawClient.APIFsGetGet(ctx, agentapi.APIFsGetGetParams{FileID: fileID})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentapi.APIFsGetGetOK:
		return typedRes.Data, nil

	case *agentapi.APIFsGetGetNotFound:
		return nil, fmt.Errorf("%w: %s", core.FileNotFoundError, typedRes.Details.Value)

	case *agentapi.APIFsGetGetBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIFsGetGetUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIFsGetGetForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIFsGetGetInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) State(ctx context.Context, includeFileIDs, includeFileSizes bool) (fsmodel.FSState, error) {
	res, err := a.rawClient.APIFsInfoPost(ctx, &agentapi.APIFsInfoPostReq{
		IncludeFileIds:   agentapi.NewOptBool(includeFileIDs),
		IncludeFileSizes: agentapi.NewOptBool(includeFileSizes),
	})
	if err != nil {
		return fsmodel.FSState{}, err
	}

	switch typedRes := res.(type) {
	case *agentapi.APIFsInfoPostOK:
		return fsmodel.FSState{
			FileIDs: typedRes.FileIds,
			Files: pkg.Map(typedRes.Files, func(raw agentapi.APIFsInfoPostOKFilesItem) fsmodel.FSStateFile {
				return fsmodel.FSStateFile{
					ID:        raw.ID,
					Size:      raw.Size,
					CreatedAt: raw.CreatedAt,
				}
			}),
			TotalFileCount: typedRes.TotalFileCount.Value,
			TotalFileSize:  typedRes.TotalFileSize.Value,
			AvailableSize:  typedRes.AvailableSize.Value,
		}, nil

	case *agentapi.APIFsInfoPostBadRequest:
		return fsmodel.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIFsInfoPostUnauthorized:
		return fsmodel.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIFsInfoPostForbidden:
		return fsmodel.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIFsInfoPostInternalServerError:
		return fsmodel.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return fsmodel.FSState{}, agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) CreateHighwayToken(ctx context.Context) (string, time.Time, error) {
	res, err := a.rawClient.APIHighwayTokenCreatePost(ctx)
	if err != nil {
		return "", time.Time{}, err
	}

	switch typedRes := res.(type) {
	case *agentapi.APIHighwayTokenCreatePostOK:
		return typedRes.Token, typedRes.ValidUntil, nil

	case *agentapi.APIHighwayTokenCreatePostBadRequest:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIHighwayTokenCreatePostUnauthorized:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIHighwayTokenCreatePostForbidden:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIHighwayTokenCreatePostInternalServerError:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return "", time.Time{}, agentmodel.AgentAPIUnknownResponse
	}
}
