package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelGetPost(
	ctx context.Context,
	req *serverapi.APILabelGetPostReq,
) (*serverapi.APILabelGetPostOK, error) {
	labels, err := c.labelUseCases.Labels(ctx, req.BookID)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
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
