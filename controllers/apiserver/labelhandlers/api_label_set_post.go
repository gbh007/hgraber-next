package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelSetPost(
	ctx context.Context,
	req *serverapi.APILabelSetPostReq,
) error {
	return c.labelUseCases.SetLabel(ctx, core.BookLabel{ //nolint:wrapcheck // будет исправлено позднее
		BookID:     req.BookID,
		PageNumber: req.PageNumber.Value,
		Name:       req.Name,
		Value:      req.Value,
	})
}
