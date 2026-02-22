package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetUpdatePost(
	ctx context.Context,
	req *serverapi.APILabelPresetUpdatePostReq,
) error {
	return c.labelUseCases.UpdateLabelPreset(ctx, core.BookLabelPreset{ //nolint:wrapcheck // будет исправлено позднее
		Name:        req.Name,
		Values:      req.Values,
		Description: req.Description.Value,
	})
}
