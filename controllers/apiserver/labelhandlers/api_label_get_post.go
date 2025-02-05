package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *LabelHandlersController) APILabelGetPost(ctx context.Context, req *serverAPI.APILabelGetPostReq) (serverAPI.APILabelGetPostRes, error) {
	labels, err := c.webAPIUseCases.Labels(ctx, req.BookID)
	if err != nil {
		return &serverAPI.APILabelGetPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelGetPostOK{
		Labels: pkg.Map(labels, func(raw core.BookLabel) serverAPI.APILabelGetPostOKLabelsItem {
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
