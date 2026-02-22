package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelDeletePost(
	ctx context.Context,
	req *serverapi.APILabelDeletePostReq,
) error {
	return c.labelUseCases.DeleteLabel(ctx, core.BookLabel{
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
	})
}
