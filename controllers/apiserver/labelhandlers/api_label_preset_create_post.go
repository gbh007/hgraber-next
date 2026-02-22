package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetCreatePost(
	ctx context.Context,
	req *serverapi.APILabelPresetCreatePostReq,
) error {
	return c.labelUseCases.CreateLabelPreset(ctx, core.BookLabelPreset{
		Name:        req.Name,
		Values:      req.Values,
		Description: req.Description.Value,
	})
}
