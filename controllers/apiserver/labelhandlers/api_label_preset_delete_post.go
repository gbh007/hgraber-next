package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetDeletePost(
	ctx context.Context,
	req *serverapi.APILabelPresetDeletePostReq,
) error {
	return c.labelUseCases.DeleteLabelPreset(ctx, req.Name)
}
