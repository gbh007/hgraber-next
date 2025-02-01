package adapter

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/agentAPI"
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
		return fmt.Errorf("%w: %s", entities.AgentAPIConflict, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostBadRequest:
		return fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostUnauthorized:
		return fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostForbidden:
		return fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsCreatePostInternalServerError:
		return fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentAPIUnknownResponse
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
		return fmt.Errorf("%w: %s", entities.FileNotFoundError, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostBadRequest:
		return fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostUnauthorized:
		return fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostForbidden:
		return fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsDeletePostInternalServerError:
		return fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentAPIUnknownResponse
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
		return nil, fmt.Errorf("%w: %s", entities.FileNotFoundError, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsGetGetInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) State(ctx context.Context, includeFileIDs, includeFileSizes bool) (entities.FSState, error) {
	res, err := a.rawClient.APIFsInfoPost(ctx, &agentAPI.APIFsInfoPostReq{
		IncludeFileIds:   agentAPI.NewOptBool(includeFileIDs),
		IncludeFileSizes: agentAPI.NewOptBool(includeFileSizes),
	})
	if err != nil {
		return entities.FSState{}, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APIFsInfoPostOK:
		return entities.FSState{
			FileIDs: typedRes.FileIds,
			Files: pkg.Map(typedRes.Files, func(raw agentAPI.APIFsInfoPostOKFilesItem) entities.FSStateFile {
				return entities.FSStateFile{
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
		return entities.FSState{}, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIFsInfoPostUnauthorized:
		return entities.FSState{}, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIFsInfoPostForbidden:
		return entities.FSState{}, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIFsInfoPostInternalServerError:
		return entities.FSState{}, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.FSState{}, entities.AgentAPIUnknownResponse
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
		return "", time.Time{}, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIHighwayTokenCreatePostUnauthorized:
		return "", time.Time{}, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIHighwayTokenCreatePostForbidden:
		return "", time.Time{}, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIHighwayTokenCreatePostInternalServerError:
		return "", time.Time{}, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return "", time.Time{}, entities.AgentAPIUnknownResponse
	}
}
