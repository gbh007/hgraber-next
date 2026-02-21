package adapter

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *FSAdapter) Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error {
	err := a.rawClient.APIFsCreatePost(
		ctx,
		agentapi.APIFsCreatePostReq{Data: body},
		agentapi.APIFsCreatePostParams{FileID: fileID},
	)
	if err != nil {
		return enrichError(err)
	}

	return nil
}

func (a *FSAdapter) Delete(ctx context.Context, fileID uuid.UUID) error {
	err := a.rawClient.APIFsDeletePost(ctx, &agentapi.APIFsDeletePostReq{FileID: fileID})
	if err != nil {
		return enrichError(err)
	}

	return nil
}

func (a *FSAdapter) Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error) {
	res, err := a.rawClient.APIFsGetGet(ctx, agentapi.APIFsGetGetParams{FileID: fileID})
	if err != nil {
		return nil, enrichError(err)
	}

	return res.Data, nil
}

func (a *FSAdapter) State(ctx context.Context, includeFileIDs, includeFileSizes bool) (fsmodel.FSState, error) {
	res, err := a.rawClient.APIFsInfoPost(ctx, &agentapi.APIFsInfoPostReq{
		IncludeFileIds:   agentapi.NewOptBool(includeFileIDs),
		IncludeFileSizes: agentapi.NewOptBool(includeFileSizes),
	})
	if err != nil {
		return fsmodel.FSState{}, enrichError(err)
	}

	return fsmodel.FSState{
		FileIDs: res.FileIds,
		Files: pkg.Map(res.Files, func(raw agentapi.APIFsInfoPostOKFilesItem) fsmodel.FSStateFile {
			return fsmodel.FSStateFile{
				ID:        raw.ID,
				Size:      raw.Size,
				CreatedAt: raw.CreatedAt,
			}
		}),
		TotalFileCount: res.TotalFileCount.Value,
		TotalFileSize:  res.TotalFileSize.Value,
		AvailableSize:  res.AvailableSize.Value,
	}, nil
}

func (a *FSAdapter) CreateHighwayToken(ctx context.Context) (string, time.Time, error) {
	res, err := a.rawClient.APIHighwayTokenCreatePost(ctx)
	if err != nil {
		return "", time.Time{}, enrichError(err)
	}

	return res.Token, res.ValidUntil, nil
}
