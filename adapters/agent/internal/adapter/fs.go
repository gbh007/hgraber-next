package adapter

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *FSAdapter) Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error {
	res, err := a.rawClient.APIFsCreatePost(
		ctx,
		agentAPI.APIFsCreatePostReq{Data: body},
		agentAPI.APIFsCreatePostParams{FileID: fileID},
	)
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIFsCreatePostNoContent:
		return nil

	case *agentAPI.APIFsCreatePostConflict:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIConflict, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostBadRequest:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostUnauthorized:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostForbidden:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostInternalServerError:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) Delete(ctx context.Context, fileID uuid.UUID) error {
	res, err := a.rawClient.APIFsDeletePost(ctx, &agentAPI.APIFsDeletePostReq{FileID: fileID})
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIFsDeletePostNoContent:
		return nil

	case *agentAPI.APIFsDeletePostNotFound:
		return fmt.Errorf("%w: %s", core.FileNotFoundError, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostBadRequest:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostUnauthorized:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostForbidden:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostInternalServerError:
		return fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	res, err := a.rawClient.APIFsGetGet(ctx, agentAPI.APIFsGetGetParams{FileID: fileID})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIFsGetGetOK:
		return typedRes.Data, nil

	case *agentAPI.APIFsGetGetNotFound:
		return nil, fmt.Errorf("%w: %s", core.FileNotFoundError, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) State(ctx context.Context, includeFileIDs, includeFileSizes bool) (core.FSState, error) {
	res, err := a.rawClient.APIFsInfoPost(ctx, &agentAPI.APIFsInfoPostReq{
		IncludeFileIds:   agentAPI.NewOptBool(includeFileIDs),
		IncludeFileSizes: agentAPI.NewOptBool(includeFileSizes),
	})
	if err != nil {
		return core.FSState{}, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIFsInfoPostOK:
		return core.FSState{
			FileIDs: typedRes.FileIds,
			Files: pkg.Map(typedRes.Files, func(raw agentAPI.APIFsInfoPostOKFilesItem) core.FSStateFile {
				return core.FSStateFile{
					ID:        raw.ID,
					Size:      raw.Size,
					CreatedAt: raw.CreatedAt,
				}
			}),
			TotalFileCount: typedRes.TotalFileCount.Value,
			TotalFileSize:  typedRes.TotalFileSize.Value,
			AvailableSize:  typedRes.AvailableSize.Value,
		}, nil

	case *agentAPI.APIFsInfoPostBadRequest:
		return core.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsInfoPostUnauthorized:
		return core.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsInfoPostForbidden:
		return core.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsInfoPostInternalServerError:
		return core.FSState{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return core.FSState{}, agentmodel.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) CreateHighwayToken(ctx context.Context) (string, time.Time, error) {
	res, err := a.rawClient.APIHighwayTokenCreatePost(ctx)
	if err != nil {
		return "", time.Time{}, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIHighwayTokenCreatePostOK:
		return typedRes.Token, typedRes.ValidUntil, nil

	case *agentAPI.APIHighwayTokenCreatePostBadRequest:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIHighwayTokenCreatePostUnauthorized:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIHighwayTokenCreatePostForbidden:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIHighwayTokenCreatePostInternalServerError:
		return "", time.Time{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return "", time.Time{}, agentmodel.AgentAPIUnknownResponse
	}
}
