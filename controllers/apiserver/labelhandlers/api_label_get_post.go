package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *LabelHandlersController) APILabelGetPost(
	ctx context.Context,
	req *serverapi.APILabelGetPostReq,
) (serverapi.APILabelGetPostRes, error) {
	labels, err := c.labelUseCases.Labels(ctx, req.BookID)
	if err != nil {
		return &serverapi.APILabelGetPostInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APILabelGetPostOK{
		Labels: pkg.Map(labels, func(raw core.BookLabel) serverapi.APILabelGetPostOKLabelsItem {
			return serverapi.APILabelGetPostOKLabelsItem{
				BookID:     raw.BookID,
				PageNumber: raw.PageNumber,
				Name:       raw.Name,
				Value:      raw.Value,
				CreatedAt:  raw.CreateAt,
			}
		}),
	}, nil
}
