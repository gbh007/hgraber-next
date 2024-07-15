package adapter

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"

	"hgnext/internal/adapters/agent/internal/client"
	"hgnext/internal/entities"
)

func (a *FSAdapter) Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error {
	res, err := a.rawClient.APIFsCreatePost(
		ctx,
		client.APIFsCreatePostReq{Data: body},
		client.APIFsCreatePostParams{FileID: fileID},
	)
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *client.APIFsCreatePostNoContent:
		return nil

	case *client.APIFsCreatePostConflict:
		return fmt.Errorf("%w: %s", entities.AgentAPIConflict, typedRes.Details.Value)

	case *client.APIFsCreatePostBadRequest:
		return fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIFsCreatePostUnauthorized:
		return fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIFsCreatePostForbidden:
		return fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIFsCreatePostInternalServerError:
		return fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) Delete(ctx context.Context, fileID uuid.UUID) error {
	res, err := a.rawClient.APIFsDeletePost(ctx, &client.APIFsDeletePostReq{FileID: fileID})
	if err != nil {
		return err
	}

	switch typedRes := res.(type) {
	case *client.APIFsDeletePostNoContent:
		return nil

	case *client.APIFsDeletePostNotFound:
		return fmt.Errorf("%w: %s", entities.FileNotFoundError, typedRes.Details.Value)

	case *client.APIFsDeletePostBadRequest:
		return fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIFsDeletePostUnauthorized:
		return fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIFsDeletePostForbidden:
		return fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIFsDeletePostInternalServerError:
		return fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	res, err := a.rawClient.APIFsGetGet(ctx, client.APIFsGetGetParams{FileID: fileID})
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *client.APIFsGetGetOK:
		return typedRes.Data, nil

	case *client.APIFsGetGetNotFound:
		return nil, fmt.Errorf("%w: %s", entities.FileNotFoundError, typedRes.Details.Value)

	case *client.APIFsGetGetBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIFsGetGetUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIFsGetGetForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIFsGetGetInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}
}

func (a *FSAdapter) IDs(ctx context.Context) ([]uuid.UUID, error) {
	res, err := a.rawClient.APIFsIdsGet(ctx)
	if err != nil {
		return nil, err
	}

	switch typedRes := res.(type) {
	case *client.APIFsIdsGetOKApplicationJSON:
		return *typedRes, nil

	case *client.APIFsIdsGetBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APIFsIdsGetUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APIFsIdsGetForbidden:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APIFsIdsGetInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.AgentAPIUnknownResponse
	}
}
