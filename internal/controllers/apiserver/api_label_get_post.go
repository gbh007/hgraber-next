package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APILabelGetPost(ctx context.Context, req *serverAPI.APILabelGetPostReq) (serverAPI.APILabelGetPostRes, error) {
	labels, err := c.webAPIUseCases.Labels(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APILabelGetPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelGetPostOK{
		Labels: pkg.Map(labels, func(raw entities.BookLabel) serverAPI.APILabelGetPostOKLabelsItem {
			return serverAPI.APILabelGetPostOKLabelsItem{
				BookID:     raw.BookID,
				PageNumber: raw.PageNumber,
				Name:       raw.Name,
				Value:      raw.Value,
				CreatedAt:  raw.CreateAt,
			}
		}),
	}, nil
}
